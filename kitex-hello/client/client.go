package main

import (
	"context"
	"errors"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	"learn-go/kitex-hello/discovery/custom"
	"learn-go/kitex-hello/kitex_gen/pbapi"
	"learn-go/kitex-hello/kitex_gen/pbapi/echo"
	"strconv"
	"sync"
	"time"
)

var addrStrs = []string{":8888", ":9999"}

func main() {
	lb := loadbalance.NewWeightedBalancer()
	//提高服务调用成功率, 有最大重试次数, 重试次数backoff等等配置
	//retryPolicy := retry.NewFailurePolicy()
	//减少请求的网络波动影响, 即一段时间没有得到返回, 会额外进行rpc请求, 有重试延迟时间,最大重试次数等等配置
	//backupPolicy := retry.NewBackupPolicy(50)
	client, err := echo.NewClient("echo",
		//client.WithHostPorts(addrStrs...),
		//client.WithFailureRetry(retryPolicy),
		//client.WithMiddleware(newFailureMW()),
		//client.WithBackupRequest(backupPolicy),
		//client.WithMiddleware(newDelayMW((60 * time.Millisecond))),
		client.WithLoadBalancer(lb),
		client.WithResolver(custom.Resolver()),
		//rpc timeout会重新resolve service instance address
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(30*time.Second),
	)
	if err != nil {
		klog.Fatal(err)
	}

	syncEcho(client, 20)

	//asyncEcho(client, 5)
}

//同步调用
func syncEcho(client echo.Client, times int) {
	for i := 0; i < times; i++ {
		req := &pbapi.Request{Message: "hello" + strconv.Itoa(i)}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			klog.Infof("fatal error: %v\n", err)
		}
		klog.Infof("return: %v\n", resp)
	}
}

//异步调用
func newFuture(f func() (any, error), callback func(any, error)) func() (any, error) {
	//rpc call返回存储
	var res any
	var err error

	//创建等待channel
	c := make(chan struct{}, 1)
	go func() {
		//rpc call会关闭channel
		defer close(c)
		//rpc call
		res, err = f()
		//callback回调
		callback(res, err)
	}()

	return func() (any, error) {
		//等待channel关闭
		<-c
		//阻塞返回
		return res, err
	}
}

func asyncEcho(client echo.Client, times int) {
	var futures []func() (any, error)
	wg := sync.WaitGroup{}
	//async rpc call
	for i := 0; i < times; i++ {
		wg.Add(1)
		var req = &pbapi.Request{Message: "async hello" + strconv.Itoa(i)}

		futures = append(futures,
			newFuture(
				func() (any, error) {
					return client.Echo(context.Background(), req)
				},
				func(resp any, err error) {
					if err != nil {
						klog.Fatalf("callback log fatal error: %v\n", err)
					}
					klog.Infof("callback return: %v\n", resp)
					wg.Done()
				}))
	}

	for i := 0; i < times; i++ {
		//阻塞等待返回
		resp, err := futures[i]()
		if err != nil {
			klog.Fatalf("block log fatal error: %v\n", err)
		}
		klog.Infof("block return: %v\n", resp)
	}

	wg.Wait()
}

//middleware, 相当于interceptor
func newFailureMW() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp any) (err error) {
			//从context中取出重试key, 如果有则是重试请求, 没有则是正常请求
			if _, exist := metainfo.GetPersistentValue(ctx, retry.TransitKey); !exist {
				println("you shall not pass")
				return kerrors.ErrRPCTimeout.WithCause(errors.New("you shall not pass"))
			}
			klog.Infof("this is a retry request")
			return next(ctx, req, resp)
		}
	}
}

func newDelayMW(delay time.Duration) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp any) (err error) {
			//从context中取出重试key, 如果有则是重试请求, 没有则是正常请求
			if _, exist := metainfo.GetPersistentValue(ctx, retry.TransitKey); !exist {
				time.Sleep(delay + 10*time.Millisecond)
				println("you shall not pass")
				return next(ctx, req, resp)
			}
			klog.Infof("this is a retry request")
			return next(ctx, req, resp)
		}
	}
}
