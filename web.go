package infra

//用于web api 的注册
var apiInitializerRegister *InitializeRegister = new(InitializeRegister)

//用户全局的注册函数
//注册API初始化对象
func RegisterApi(ai Initializer) {
	apiInitializerRegister.Register(ai)
}

//获取注册的web 挨批出书啊对象
func GetApiInitializers() []Initializer {
	return apiInitializerRegister.Initializers
}

type WebApiStart struct {
	BaseStarter
}

//需要在setup阶段来运行这些初始化
func (w *WebApiStart) Setup(ctx StarterContext) {
	for _, v := range GetApiInitializers() {
		v.Init()
	}
}
