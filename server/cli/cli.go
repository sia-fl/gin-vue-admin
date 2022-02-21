package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/cli/resolve"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "fresh:db",
				Usage: "刷新数据库",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "type",
						Usage: "使用 postgres、mysql 作为数据库源",
					},
				},
				Action: func(context *cli.Context) error {
					dbType := context.String("type")
					if dbType == "" {
						dbType = "mysql"
					}
					return resolve.FreshDb(dbType)
				},
			},
		},
	}
	global.GVA_VP = core.Viper("./config.yaml") // 初始化Viper
	global.GVA_DB = initialize.Gorm()           // gorm连接数据库
	err := app.Run(os.Args)
	if err != nil {
		print(err.Error())
	}
}
