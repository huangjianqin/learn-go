package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"learn-go/grpc-hello/proto"
)

const (
	defaultName = "world"
)

var (
	//addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

//定义服务名
const (
	exampleScheme      = "helloword"
	exampleServiceName = "grpc.helloword"
)

//定义服务upstream地址
var (
	addrs         = []string{"localhost:50051", "localhost:50052"}
	serviceConfig = `{
		"loadBalancingConfig": [{"round_robin":{}}],
		"methodConfig": [
			{
			  "name": [{"service": "helloworld.Greeter"}],
			  "waitForReady": true,
			  "retryPolicy": {
				  "MaxAttempts": 4,
				  "InitialBackoff": ".01s",
				  "MaxBackoff": ".01s",
				  "BackoffMultiplier": 1.0,
				  "RetryableStatusCodes": [ "UNAVAILABLE" ]
			  }
			}
		]
		}`
)

//grpc-go要支持retry必须要设置GRPC_GO_RETRY=on
const RETRY_ON_ENV_KEY = "GRPC_GO_RETRY"

func main() {
	env, exists := os.LookupEnv(RETRY_ON_ENV_KEY)
	if exists {
		fmt.Printf("%v = %v\n", RETRY_ON_ENV_KEY, env)
	} else {
		fmt.Printf("%v not exists!!!\n", RETRY_ON_ENV_KEY)
	}

	//解析命令行输入并赋值到addr和name变量
	//flag.Parse()
	//创建connection
	conn, err := grpc.Dial(
		//*addr,
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig),
		grpc.WithUnaryInterceptor(unaryInterceptor()),
		grpc.WithStreamInterceptor(streamInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//构建服务client
	c := helloworld.NewGreeterClient(conn)

	reqRespN(c, 5)

	//biStream(c)
}

//req-resp
func reqRespN(c helloworld.GreeterClient, num int) {
	for i := 0; i < num; i++ {
		reqResp(c)
	}
}
func reqResp(c helloworld.GreeterClient) {
	//加上元数据
	md := metadata.Pairs("author", "cc")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//=WithDeadline, 即超过某个deadline, rpc call就会被取消
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	var header, trailer metadata.MD
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: *name}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())

	if t, ok := header["timestamp"]; ok {
		fmt.Printf("timestamp from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}

	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in trailer")
	}
}

//stream
func biStream(c helloworld.GreeterClient) {
	biClient, err := c.KeepSayHello(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}

	go func() {
		for range time.Tick(time.Second) {
			ierr := biClient.Send(&helloworld.HelloRequest{Name: *name})
			if ierr != nil {
				log.Fatalf("failed to send greet: %v", ierr)
				break
			}
		}
		biClient.CloseSend()
	}()

	for true {
		resp, err := biClient.Recv()
		if err == io.EOF {
			//stream end
			break
		}

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("recv %v\n", resp.GetMessage())
	}
}

//unaryInterceptor
func unaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		end := time.Now()
		log.Printf("RPC: %v, start time: %v, end time: %v, err: %v", method, start, end, err)
		return err
	}
}

//streamInterceptor
func streamInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		s, err := streamer(ctx, desc, cc, method, opts...)
		if err != nil {
			return nil, err
		}
		return newWrappedStream(s), nil
	}
}

type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now())
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now())
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

//--------自定义name resolver builder
type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			exampleServiceName: addrs,
		},
	}
	r.start()
	return r, nil
}
func (*exampleResolverBuilder) Scheme() string { return exampleScheme }

//--------自定义name resolver
type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*exampleResolver) Close()                                  {}

func init() {
	resolver.Register(&exampleResolverBuilder{})
}
