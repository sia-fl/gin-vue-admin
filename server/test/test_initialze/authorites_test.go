package test_initialze

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemSource "github.com/flipped-aurora/gin-vue-admin/server/source/system"
	"github.com/flipped-aurora/gin-vue-admin/server/test"
	. "github.com/go-playground/assert/v2"
	"testing"
)

func TestInitAuthorites(t *testing.T) {
	test.Start()
	orm := global.GVA_DB
	Equal(t, orm.Migrator().DropTable(system.SysBaseMenu{}), nil)
	Equal(t, orm.AutoMigrate(system.SysBaseMenu{}), nil)
	Equal(t, systemSource.BaseMenu.Initialize(), nil)
}
