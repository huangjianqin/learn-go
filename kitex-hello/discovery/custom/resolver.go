package custom

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var addrStrs = []string{":8888", ":9999"}
var resolver = &CustomResolver{}

type CustomResolver struct {
}

func Resolver() discovery.Resolver {
	return resolver
}

func (p *CustomResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	//模拟注册中心返回服务网络信息
	instances := make([]discovery.Instance, cap(addrStrs))
	for i, addrStr := range addrStrs {
		weight := rand.Intn(100) + 1
		instances[i] = NewInstance(NewNetAddr("tcp", addrStr), weight)
		//打印一下权重
		fmt.Printf("%v-%v\n", addrStr, weight)
	}

	return discovery.Result{Cacheable: true, CacheKey: desc, Instances: instances}, nil
}

//返回endpoint对应的描述信息, 用作cache key
func (p *CustomResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	//注意, target.Address()这里返回nil, 同时会阻塞, 因为我们没给它赋值
	return target.ServiceName()
}

//返回resolver name
func (p *CustomResolver) Name() string {
	return "CustomResolver"
}

//返回两次服务发现结果的不同
// Diff computes the difference between two results.
// When `next` is cacheable, the Change should be cacheable, too. And the `Result` field's CacheKey in
// the return value should be set with the given cacheKey.
func (p *CustomResolver) Diff(cacheKey string, prev discovery.Result, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}
