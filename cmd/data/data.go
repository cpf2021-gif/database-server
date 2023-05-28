package main

import (
	"server/global"
	"server/model/inventory"
	"server/model/product"
	"server/model/user"
	"server/setup"
)

// 初始数据录入
func main() {
	// 配置文件
	global.GL_VIPER = setup.InitializeViper("./")

	// 连接数据库
	global.GL_DB = setup.InitializeDB()

	// 创建表
	global.GL_DB.AutoMigrate(&user.User{}, &product.Product{}, &product.Supplier{}, &inventory.Inventory{}, &inventory.Inbound{}, &inventory.Outbound{})

	// 录入数据
	// 用户
	user.InitializeUser(global.GL_DB)

	// 产品
	product.InitializeProduct(global.GL_DB)

	// 库存
	inventory.InitializeInventory(global.GL_DB)
}
