package main

import (
	"server/global"
	"server/setup"
)

func main() {
	// 配置文件
	global.GL_VIPER = setup.InitializeViper("./")
}