package infra

import "github.com/tietang/props/kvs"

const (
	KeyProps = "_conf"
)

//基础资源结构体
type StarterContext map[string]interface{}

func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没被初始化")
	}
	//因为map中存储的是interface类型,所以将p 转换为kvs.ConfigSource
	return p.(kvs.ConfigSource)
}

//基础资源启动接口
type Starter interface {
	//1.系统启动，初始化一些基础资源
	Init(StarterContext)
	//2.系统基础资源安装
	Setup(StarterContext)
	//3.启动基础资源
	Start(StarterContext)
	//启动器是否可阻塞
	StartBlocking() bool
	//4.资源 停止和销毁
	Stop(StarterContext)
}

var _ Starter = new(BaseStarter)

//基础空启动器，为了方便资源启动器的代码实现
type BaseStarter struct {
}

func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}

//启动器注册器
type starterRegister struct {
	starters []Starter
}

//启动器注册
func (s *starterRegister) Register(ss Starter) {
	s.starters = append(s.starters, ss)
}

//获取所有
func (s *starterRegister) AllStarters() []Starter {
	return s.starters
}

var StarterRegister *starterRegister = new(starterRegister)

func Register(s Starter) {
	StarterRegister.Register(s)
}

//系统基础资源的启动管理
func SysRun() {
	//1.初始化
	ctx := StarterContext{}

	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(ctx)
	}

	//2.安装
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(ctx)
	}

	//3.启动
	for _, starter := range StarterRegister.AllStarters() {
		starter.Start(ctx)
	}

}
