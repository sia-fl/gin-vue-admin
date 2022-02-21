package resolve

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	serviceSystem "github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

func FreshDb(dbType string) error {
	initialize.DBList()
	if dbType == "" {
		dbType = "mysql"
	}
	initDb := request.NewInitDb()
	db := initDb.InitSqlDb()
	dropDbSql := fmt.Sprintf("drop database if exists %s", initDb.DBName)
	createDbSql := fmt.Sprintf("create database %s", initDb.DBName)
	_, err := db.Exec(dropDbSql)
	if err != nil {
		return err
	}
	_, err = db.Exec(createDbSql)
	if err != nil {
		return err
	}
	dbService := serviceSystem.InitDBService{}
	return dbService.InitDB(*initDb)
}
