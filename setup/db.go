package setup

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/global"
	"server/util"

	"server/model/inventory"
	"server/model/product"
	"server/model/user"
)

// 返回数据库连接
func InitializeDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(global.GL_CONFIG.Database.GetDSN()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&user.User{}, &product.Product{}, &product.Supplier{}, &inventory.Inventory{}, &inventory.Inbound{}, &inventory.Outbound{}, &product.Seller{})

	// 每次启动时备份数据库
	if err := util.Backup(global.GL_VIPER); err != nil {
		panic(err)
	}

	fmt.Println("database initialized")
	return db
}
