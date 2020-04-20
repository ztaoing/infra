package base

import (
	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func init() {
	//定义日志的格式-文本
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	formatter.ForceFormatting = true

	//根据级别设置颜色
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "",
		FatalLevelStyle: "",
		PanicLevelStyle: "",
		DebugLevelStyle: "",
		PrefixStyle:     "",
		TimestampStyle:  "38",
	})

	log.SetFormatter(formatter)

	//日志的级别-默认为info
	level := os.Getenv("log.debug")
	if level == "true" {
		log.SetLevel(log.DebugLevel)
	}

	//控制台高亮显示
	formatter.ForceColors = true
	formatter.DisableColors = false

	//日志文件和滚动配置
	log.Info("测试")
	log.Debug("debug")
}
