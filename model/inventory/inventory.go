package inventory

import (
	"time"

	"gorm.io/gorm"

	"server/global"
	"server/model/product"
	"server/model/user"
)

// 库存表
/*
CREATE TABLE inventories
(
	id           BIGSERIAL PRIMARY KEY,
	product_name varchar(20) NOT NULL,
	quantity     bigint NOT NULL,
	max_quantity bigint NOT NULL,
	min_quantity bigint NOT NULL,
	create_time  timestamp with time zone not null,
	update_time  timestamp with time zone not null,
	FOREIGN KEY (product_name) REFERENCES products(name) ON UPDATE CASCADE ON DELETE RESTRICT
);
*/

type Inventory struct {
	ID          int             `json:"id" gorm:"primary_key"`
	ProductName string          `json:"product_name" gorm:"type:varchar(20);not null"`
	Product     product.Product `json:"product" gorm:"references:name;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity    int             `json:"quantity" gorm:"type:int;not null"`
	MaxQuantity int             `json:"max_quantity" gorm:"type:int;not null"`
	MinQuantity int             `json:"min_quantity" gorm:"type:int;not null"`

	CreateTime time.Time `json:"create_time" gorm:"not null"`
	UpdateTime time.Time `json:"update_time" gorm:"not null"`
}

func (i *Inventory) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	i.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	i.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

func (i *Inventory) BeforeUpdate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	i.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

// 入库表
/*
CREATE TABLE inbound
(
	id          BIGSERIAL PRIMARY KEY,
	product_name varchar(20) NOT NULL,
	quantity    bigint NOT NULL,
	user_name   varchar(20) NOT NULL,
	create_time timestamp with time zone not null,
	FOREIGN KEY (product_name) REFERENCES products(name) ON UPDATE CASCADE ON DELETE RESTRICT,
	FOREIGN KEY (user_name) REFERENCES users(username) ON UPDATE CASCADE ON DELETE RESTRICT
);
*/

type Inbound struct {
	ID int `json:"id" gorm:"primary_key"`

	ProductName string          `json:"product_name" gorm:"type:varchar(20);not null"`
	Product     product.Product `json:"product" gorm:"references:name;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Quantity int `json:"quantity" gorm:"type:int;not null"`

	UserName string    `json:"user_name" gorm:"type:varchar(20);not null"`
	User     user.User `json:"user" gorm:"foreignkey:UserName;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	CreateTime time.Time `json:"create_time" gorm:"index; not null"`
}

func (i *Inbound) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	i.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

// 出库表
/*
CREATE TABLE outbound
(
	id          BIGSERIAL PRIMARY KEY,
	product_name varchar(20) NOT NULL,
	quantity    bigint NOT NULL,
	user_name   varchar(20) NOT NULL,
	create_time timestamp with time zone not null,
	FOREIGN KEY (product_name) REFERENCES products(name) ON UPDATE CASCADE ON DELETE RESTRICT,
	FOREIGN KEY (user_name) REFERENCES users(username) ON UPDATE CASCADE ON DELETE RESTRICT
);
*/

type Outbound struct {
	ID int `json:"id" gorm:"primary_key"`

	ProductName string          `json:"product_name" gorm:"type:varchar(20);not null"`
	Product     product.Product `json:"product" gorm:"references:name;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Quantity int `json:"quantity" gorm:"type:int;not null"`

	UserName string    `json:"user_name" gorm:"type:varchar(20);not null"`
	User     user.User `json:"user" gorm:"foreignkey:UserName;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	CreateTime time.Time `json:"create_time" gorm:"index; not null"`
}

func (o *Outbound) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	o.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

func InitializeInventory(db *gorm.DB) {
	// 第一次入库
	/*
		INSERT INTO inventory (product_name, quantity, max_quantity, min_quantity, create_time, update_time)
		VALUES ('iPhone 13', 100, 200, 50, '2021-10-01 00:00:00', '2021-10-01 00:00:00'),
	*/

	inventorys := []Inventory{
		{ProductName: "iPhone 13", Quantity: 100, MaxQuantity: 200, MinQuantity: 50},
		{ProductName: "MacBook Pro", Quantity: 70, MaxQuantity: 140, MinQuantity: 35},
		{ProductName: "Mate 50", Quantity: 50, MaxQuantity: 100, MinQuantity: 25},
		{ProductName: "HuaWei Watch 3", Quantity: 60, MaxQuantity: 120, MinQuantity: 30},
		{ProductName: "HuaWei MatePad Pro", Quantity: 80, MaxQuantity: 160, MinQuantity: 40},
		{ProductName: "Galaxy S22", Quantity: 40, MaxQuantity: 80, MinQuantity: 20},
		{ProductName: "Surface Pro 8", Quantity: 30, MaxQuantity: 60, MinQuantity: 15},
		{ProductName: "Pixel 6", Quantity: 15, MaxQuantity: 30, MinQuantity: 7},
		{ProductName: "PlayStation 5", Quantity: 90, MaxQuantity: 180, MinQuantity: 45},
		{ProductName: "Kindle", Quantity: 78, MaxQuantity: 156, MinQuantity: 39},
		{ProductName: "Switch", Quantity: 66, MaxQuantity: 132, MinQuantity: 33},
		{ProductName: "Xiaomi 12", Quantity: 88, MaxQuantity: 176, MinQuantity: 44},
		{ProductName: "Xiaomi Pad 5", Quantity: 99, MaxQuantity: 198, MinQuantity: 49},
		{ProductName: "Xiaomi Watch", Quantity: 100, MaxQuantity: 200, MinQuantity: 50},
		{ProductName: "OPPO Find X5", Quantity: 30, MaxQuantity: 60, MinQuantity: 15},
		{ProductName: "OPPO Watch 2", Quantity: 68, MaxQuantity: 136, MinQuantity: 34},
	}

	global.GL_DB.Model(&Inventory{}).Create(&inventorys)

	for _, inventory := range inventorys {
		// 第一次入库
		/*
			INSERT INTO inbound (product_name, quantity, user_name, create_time)
			VALUES ('iPhone 13', 100, 'admin', '2021-10-01 00:00:00'),
		*/
		inbound := Inbound{
			ProductName: inventory.ProductName,
			Quantity:    inventory.Quantity,
			UserName:    "admin",
		}
		global.GL_DB.Model(&Inbound{}).Create(&inbound)
	}
}
