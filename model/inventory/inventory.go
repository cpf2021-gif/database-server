package inventory

import (
	"time"

	"gorm.io/gorm"

	"server/model/product"
	"server/model/user"
)

// 库存表
type Inventory struct {
	ID          int             `json:"id" gorm:"primary_key"`
	ProductID   int             `json:"product_id" gorm:"type:int;not null"`
	Product     product.Product `json:"product" gorm:"foreignkey:ProductID"`
	Quantity    int             `json:"quantity" gorm:"type:int;not null"`
	MaxQuantity int             `json:"max_quantity" gorm:"type:int;not null"`
	MinQuantity int             `json:"min_quantity" gorm:"type:int;not null"`

	CreateTime time.Time `json:"create_time" gorm:"index"`
	UpdateTime time.Time `json:"update_time" gorm:"index"`
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
type Inbound struct {
	ID int `json:"id" gorm:"primary_key"`

	ProductID int             `json:"product_id" gorm:"type:int;not null"`
	Product   product.Product `json:"product" gorm:"foreignkey:ProductID"`

	SupplierID int              `json:"supplier_id" gorm:"type:int;not null"`
	Supplier   product.Supplier `json:"supplier" gorm:"foreignkey:SupplierID"`

	Quantity int `json:"quantity" gorm:"type:int;not null"`

	UserID int       `json:"user_id" gorm:"type:int;not null"`
	User   user.User `json:"user" gorm:"foreignkey:UserID"`

	CreateTime time.Time `json:"create_time" gorm:"index"`
}

func (i *Inbound) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	i.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

// 出库表
type Outbound struct {
	ID int `json:"id" gorm:"primary_key"`

	ProductID int             `json:"product_id" gorm:"type:int;not null"`
	Product   product.Product `json:"product" gorm:"foreignkey:ProductID"`

	Quantity int `json:"quantity" gorm:"type:int;not null"`

	UserID int       `json:"user_id" gorm:"type:int;not null"`
	User   user.User `json:"user" gorm:"foreignkey:UserID"`

	CreateTime time.Time `json:"create_time" gorm:"index"`
}

func (o *Outbound) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	o.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}
