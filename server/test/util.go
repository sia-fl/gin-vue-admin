package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"sync"
)

var initLock = &sync.Once{}

func Start() {
	initLock.Do(func() {
		global.GVA_VP = core.Viper("../../config.yaml") // 初始化Viper
		global.GVA_DB = initialize.Gorm()               // gorm连接数据库
		initialize.DBList()
	})
}
