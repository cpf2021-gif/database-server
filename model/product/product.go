package product

import (
	"time"

	"gorm.io/gorm"
	"server/model/supplier"
)

// 产品表
type Product struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Name       string    `json:"name" gorm:"type:varchar(20);not null"`
	CreateTime time.Time `json:"create_time" gorm:"index"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	p.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

// 产品和供应商表的中间表
type ProductSupplier struct {
	ID         int               `json:"id" gorm:"primary_key"`
	ProductID  int               `json:"product_id" gorm:"type:int;not null"`
	Product    Product           `json:"product" gorm:"foreignkey:ProductID"`
	SupplierID int               `json:"supplier_id" gorm:"type:int;not null"`
	Supplier   supplier.Supplier `json:"supplier" gorm:"foreignkey:SupplierID"`

	CreateTime time.Time `json:"create_time" gorm:"index"`
}

func (ps *ProductSupplier) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	ps.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}
