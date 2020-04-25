package gorpc

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/ztaoing/infra/lb"
	"net/rpc"
)

type GoRpcClient struct {
	apps *lb.Apps
}

//基于服务发现和负载均衡实现的rpc
func (g *GoRpcClient) Call(serviceId, serviceMethod string, in interface{}, out interface{}) error {
	//通过微服务名称从被年底服务注册表中查询应用和应用实例
	app := g.apps.Get(serviceId)
	if app == nil {
		return errors.New("没有可用的微服务应用，应用名称：" + serviceId + ",请求：" + serviceMethod)
	}
	//通过负载均衡算法从应用实例列表中算法一个实例
	ins := app.Get(serviceMethod)
	if ins == nil {
		return errors.New("没有可用的应用实例，应用名称：" + serviceId + ",请求：" + serviceMethod)
	}
	//通过负载均衡算法获得的实例的ip和端口
	addr := ins.Metadata["rpcAddr"]
	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Error(err)
		return err
	}
	defer c.Close()
	//链接建立成功
	err = c.Call(serviceMethod, in, out)
	if err != nil {
		log.Error(err)
		return err
	}
	return err

}
