package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"learn-go/grpc-hello/proto"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var addrs = []string{":50051", ":50052"}

// 实现 helloworld.GreeterServer
type server struct {
	helloworld.UnimplementedGreeterServer
	addr    string
	counter int32
}

// 实现helloworld.HelloRequest消息处理逻辑
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received on %v: %v", s.addr, in.GetName())

	//创建trailer
	defer func() {
		trailer := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
		grpc.SetTrailer(ctx, trailer)
	}()

	//读取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}

	fmt.Printf("read from metadata:\n")
	for i, data := range md {
		fmt.Printf(" %v : %v\n", i, data)
	}

	// Create and send header.
	header := metadata.New(map[string]string{"location": "MTV", "timestamp": time.Now().Format(time.StampNano)})
	grpc.SendHeader(ctx, header)

	afterCount := atomic.AddInt32(&s.counter, 1)
	if afterCount%2 == 0 {
		//模拟失败
		//返回unavailable code, client可以配置哪些code允许重试
		return nil, status.Errorf(codes.Unavailable, "fail")
	}

	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) KeepSayHello(stream helloworld.Greeter_KeepSayHelloServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			fmt.Printf("server: error receiving from stream: %v\n", err)
			return err
		}
		log.Printf("Received: %v\n", in.GetName())
		stream.Send(&helloworld.HelloReply{Message: "Hello " + in.GetName()})
	}
}

func main() {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			startServer(addr)
		}(addr)
	}
	wg.Wait()
}

func startServer(addr string) {
	//监听端口
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//new grpc server
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor()), grpc.StreamInterceptor(streamInterceptor()))

	//注册服务
	//health服务
	//healthcheck := health.NewServer()
	//healthgrpc.RegisterHealthServer(s, healthcheck)
	//根据条件设置health状态
	//healthcheck.SetServingStatus("helloworld.Greeter", healthgrpc.HealthCheckResponse_NOT_SERVING)

	//Greeter服务
	helloworld.RegisterGreeterServer(s, &server{addr: addr})
	log.Printf("server listening at %v", lis.Addr())
	//暴露服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func unaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		m, err := handler(ctx, req)
		if err != nil {
			log.Printf("RPC failed with error %v", err)
		}
		log.Printf("Send a message (Type: %T) at %v", m, time.Now())
		return m, err
	}
}

//ServerStream封装
type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now())
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now())
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func streamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, newWrappedStream(ss))
		if err != nil {
			log.Printf("RPC failed with error %v", err)
		}
		return err
	}
}
