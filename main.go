package main

import (
	"server/global"
	"server/routers"
	"server/setup"
)

func main() {
	// 配置文件
	global.GL_VIPER = setup.InitializeViper("./")

	// 连接数据库
	global.GL_DB = setup.InitializeDB()

	// 初始化路由
	r := routers.InitRouter()
	r.Run(":8080")
}