package setup

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/global"
	"server/model/user"
)

func InitializeDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(global.GL_CONFIG.Database.GetDSN()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	fmt.Println("database initialized")

	// Migrate the schema
	db.AutoMigrate(&user.User{})

	return db
}
