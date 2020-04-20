package base

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisrecover "github.com/kataras/iris/middleware/recover"
	log "github.com/sirupsen/logrus"
	"go1234.cn/newResk/infra"
	"time"
)

var irisApplication *iris.Application

//创建iris实例
func Iris() *iris.Application {
	Check(irisApplication)
	return irisApplication
}

type IrisSveverStarter struct {
	infra.BaseStarter
}

//iris需要在 Init和Start阶段处理
func (i *IrisSveverStarter) Init(ctx infra.StarterContext) {
	//在此阶段需要做以下事情
	//创建iris application实例
	irisApplication := initIris()
	//日志组件配置和扩展
	logger := irisApplication.Logger()
	logger.Install(log.StandardLogger())
	//主要中间件的配置：recover，日志输出中间件的自定义

}

func (i *IrisSveverStarter) Start(ctx infra.StarterContext) {
	//与logrus保持一致的日志级别
	Iris().Logger().SetLevel(ctx.Props().GetDefault("log.level", "info"))
	//把路由信息打印到控制台
	routes := Iris().GetRoutes()
	for _, v := range routes {
		log.Info(v.Trace())
	}
	//启动iris
	port := ctx.Props().GetDefault("app.server.port", "18080")
	Iris().Run(iris.Addr(":" + port))
}

func initIris() *iris.Application {
	app := iris.New()
	app.Use(irisrecover.New())
	//中间件的配置
	cfg := logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		Columns:            true,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			app.Logger().Info("| %s | %s | %s | %s | %s | %s | %s | %s", now.Format("2006-01-02.15.04.05.000000"),
				latency.String(), status, ip, method, path, message, headerMessage)
		},
		Skippers: nil,
	}
	app.Use(logger.New(cfg))
	return app
}
