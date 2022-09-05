package custom

import (
	"net"

	"github.com/cloudwego/kitex/pkg/discovery"
)

//实现discovery.Instance接口
type Instance struct {
	netAddr net.Addr
	weight  int
}

func NewInstance(netAddr net.Addr, weight int) discovery.Instance {
	return &Instance{
		netAddr: netAddr,
		weight:  weight,
	}
}

func (i *Instance) Address() net.Addr {
	return i.netAddr
}

func (i *Instance) Weight() int {
	return i.weight
}

func (i *Instance) Tag(key string) (value string, exist bool) {
	return "", false
}

//net.Addr接口
type NetAddr struct {
	network string
	address string
}

func NewNetAddr(network, address string) net.Addr {
	return &NetAddr{network, address}
}

func (na *NetAddr) Network() string {
	return na.network
}

func (na *NetAddr) String() string {
	return na.address
}
