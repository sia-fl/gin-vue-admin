package test_cli

import (
	"github.com/flipped-aurora/gin-vue-admin/server/cli/resolve"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/test"
	. "github.com/go-playground/assert/v2"
	"testing"
)

func TestFreshDb(t *testing.T) {
	test.Start()
	dbType := "mysql"
	if global.GVA_CONFIG.Pgsql.Path != "" {
		dbType = "postgres"
	}
	Equal(t, resolve.FreshDb(dbType), nil)
}
