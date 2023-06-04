package main

import (
	"time"

	"github.com/go-co-op/gocron"

	"server/global"
	"server/routers"
	"server/setup"
	"server/util"
)

func main() {
	// 配置文件
	global.GL_VIPER = setup.InitializeViper("./")

	// 连接数据库
	global.GL_DB = setup.InitializeDB()

	// 定期备份数据库
	s := gocron.NewScheduler(time.Local)
	s.Every(10).Minutes().Do(func() {
		util.Backup(global.GL_VIPER)
	})

	s.StartAsync()

	// 初始化路由
	r := routers.InitRouter()
	r.Run(":8080")
}
