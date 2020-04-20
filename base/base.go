package base

import (
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

//检查传入的interface是否为nil，如果为nil，就通过log.Panic函数抛出异常
//用在starter中检查公共资源是否被实例化
func Check(a interface{}) {
	if a == nil {
		_, file, l, _ := runtime.Caller(1)
		strs := strings.Split(file, "/")
		size := len(strs)
		if size > 4 {
			size = 4
		}
		file = filepath.Join(strs[len(strs)-size:]...)
		log.Panicf("对象不能为空，%s(%d)", file, l)
	}
}
