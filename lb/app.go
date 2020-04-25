package lb

import (
	"fmt"
	"github.com/tietang/go-eureka-client/eureka"
	"strings"
)

//服务实例的状态
type status string

const (
	StatusEnabled  status = "enabled"
	StatusDisabled status = "disabled"
)

type ServerInstance struct {
	InstanceId string
	AppName    string
	Address    string
	Status     status
	Metadata   map[string]string
}

//用于管理所有的应用程序
type Apps struct {
	client eureka.Client
}

//通过应用程序名称在本地的应用程序缓存列表中得到应用程序
func (a *Apps) Get(appName string) *App {
	//从服务注册中心中查询应用程序
	//Applications是缓存在本地的注册列表
	var app *eureka.Application
	for _, a := range a.client.Applications.Applications {
		if a.Name == strings.ToUpper(appName) {
			app = &a
			break
		}
	}
	if app == nil {
		return nil
	}

	newA := &App{
		Name:      app.Name,
		Instances: make([]*ServerInstance, 0),
	}
	//
	for _, ins := range app.Instances {
		var port int
		if ins.SecurePort.Enabled {
			port = ins.SecurePort.Port
		} else {
			port = ins.Port.Port
		}
		//构造服务实例
		si := &ServerInstance{
			InstanceId: ins.InstanceId,
			AppName:    appName,
			Status:     status(ins.Status),
			Address:    fmt.Sprintf("%s:%d", ins.IpAddr, port),
			Metadata:   make(map[string]string),
		}
		//追加
		si.Metadata["rpcAddr"] = fmt.Sprintf("%s:%d", ins.Metadata.Map["port"], port)
		newA.Instances = append(newA.Instances, si)
	}
	return newA
}

type App struct {
	Name      string
	Instances []*ServerInstance
	lb        Balancer //每一个应用可以设置不同的算法
}

func (a *App) Get(key string) *ServerInstance {
	//通过负载均衡算法获取实例
	ins := a.lb.Next(key, a.Instances)
	return ins

}
