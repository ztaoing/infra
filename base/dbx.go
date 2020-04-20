package base

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"go1234.cn/newResk/infra"
)

//把数据库的初始化放在setup阶段
var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

//dbx的starter
type DbxStarter struct {
	infra.BaseStarter
}

func (s *DbxStarter) Setup(ctx infra.StarterContext) {
	fmt.Println("here")
	conf := ctx.Props()
	//初始化数据库配置
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	logrus.Info("mysql.conn url:", settings.ShortDataSourceName())
	dbx, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(dbx.Ping())
	database = dbx
}
