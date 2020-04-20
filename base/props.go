package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"go1234.cn/newResk/infra"
	"sync"
)

var props kvs.ConfigSource

//供外部调用
func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {

	props = ctx.Props()
	fmt.Println("初始化配置")
	GetSystemAccount()

}

//存储加载的配置
type SystemAccount struct {
	AccountNum  string
	AccountName string
	UserId      string
	Username    string
}

var systemAccount *SystemAccount
var systemAccountOnce sync.Once

//获取系统账户
func GetSystemAccount() *SystemAccount {
	systemAccountOnce.Do(func() {
		systemAccount = new(SystemAccount)
		err := kvs.Unmarshal(Props(), systemAccount, "system.account")
		if err != nil {
			//没有获得系统红包账户，则panic
			panic(err)
		}

	})
	return systemAccount
}

//获取红包链接
func GetEnvelopeActivityLink() string {
	link := Props().GetDefault("envelope.link", "/v1/envelope/link")
	return link
}

//获取域名
func GetEnvelopeDomain() string {
	domain := Props().GetDefault("envelope.domain", "http://localhost")
	return domain
}
