package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"learn-go/kitex-hello/kitex_gen/pbapi"
	"learn-go/kitex-hello/kitex_gen/pbapi/echo"
	"log"
	"math/rand"
	"net"
	"time"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Response, err error) {
	klog.Infof("echo called, received: %v", req.GetMessage())
	rand := rand.Intn(10) + 1
	time.Sleep(time.Duration(rand) * 100 * time.Millisecond)
	return &pbapi.Response{Message: "service reply >> " + req.Message}, nil
}

func main() {
	//真随机
	rand.Seed(time.Now().UnixNano())

	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	svr := echo.NewServer(&EchoImpl{}, server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
