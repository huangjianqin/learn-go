package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	"learn-go/kitex-hello/kitex_gen/pbapi"
	"learn-go/kitex-hello/kitex_gen/pbapi/echo"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct {
	addr string
}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Response, err error) {
	klog.Infof("%v >> echo called, received: %v", s.addr, req.GetMessage())
	rand := rand.Intn(10) + 1
	//timeout测试
	//time.Sleep(time.Duration(rand) * 1000 * time.Millisecond)
	time.Sleep(time.Duration(rand) * 100 * time.Millisecond)
	return &pbapi.Response{Message: "service reply >> " + req.Message}, nil
}

var addrStrs = []string{":8888", ":9999"}

func main() {
	//真随机
	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}
	wg.Add(cap(addrStrs))
	for _, addrStr := range addrStrs {
		go startServer(wg, addrStr)
	}
	wg.Wait()
}

func startServer(wg sync.WaitGroup, addrStr string) {
	defer wg.Done()

	addr, _ := net.ResolveTCPAddr("tcp", addrStr)
	server := echo.NewServer(
		&EchoImpl{addr: addrStr},
		server.WithServiceAddr(addr),
		server.WithReadWriteTimeout(500*time.Millisecond),
		server.WithMaxConnIdleTime(30*time.Second),
		//每秒5个请求
		server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 5}),
	)

	err := server.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
