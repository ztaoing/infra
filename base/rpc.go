package base

import (
	log "github.com/sirupsen/logrus"
	"go1234.cn/newResk/infra"
	"reflect"

	"net"
	"net/rpc"
)

//用rpc接口来替换web接口
var rpcServer *rpc.Server

func RpcServer() *rpc.Server {
	Check(rpcServer)

	return rpcServer
}

//注册rpcserver
func RpcRegister(ri interface{}) {
	//获取ri的类型
	typ := reflect.TypeOf(ri)
	//在控制台打印
	log.Infof("goRpc Register:%s", typ.String())

	RpcServer().Register(ri)
}

type GoRPCStarter struct {
	infra.BaseStarter
	server *rpc.Server
}

func (g *GoRPCStarter) Init(ctx infra.StarterContext) {
	g.server = rpc.NewServer()
	rpcServer = g.server
}
func (g *GoRPCStarter) Start(ctx infra.StarterContext) {

	port := ctx.Props().GetDefault("app.rpc.port", "8082")
	//监听网络端口
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panic(err)
	}
	//监听成功
	log.Info("listening tcp  for rpc with port:", port)
	//处理网络请求
	go server.Accept(listener)
	//将server注册
	rpcServer = server

}
