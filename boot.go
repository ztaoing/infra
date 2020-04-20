package infra

import "github.com/tietang/props/kvs"

/**
管理程序的启动、加载的所有生命周期
*/

type BootaApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootaApplication {
	b := &BootaApplication{
		conf:           conf,
		starterContext: StarterContext{},
	}

	//
	b.starterContext[KeyProps] = conf
	return b
}

//定义整个初始化的生命周期
func (b *BootaApplication) Start() {
	//1.初始化starter
	b.init()
	//2.安装starter
	b.setup()
	//3.启动starter
	b.start()

}

func (b *BootaApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootaApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}
func (b *BootaApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {
		if starter.StartBlocking() {
			//可阻塞的
			//如果是最后一个,直接启动
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else {
				//如果不是最后一个,异步启动，防止阻塞后边的starter
				go func() {
					starter.Start(b.starterContext)
				}()
			}

		} else {
			//不可阻塞的
			starter.Start(b.starterContext)
		}

	}
}
