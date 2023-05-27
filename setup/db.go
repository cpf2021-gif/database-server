package setup

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/global"
	"server/model/inventory"
	"server/model/product"
	"server/model/user"
)

// 初始化数据库
func InitializeDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(global.GL_CONFIG.Database.GetDSN()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	fmt.Println("database initialized")

	// Migrate the schema
	db.AutoMigrate(&user.User{}, &product.Product{}, &product.Supplier{}, &product.ProductSupplier{}, &inventory.Inventory{}, &inventory.Inbound{}, &inventory.Outbound{})

	return db
}
