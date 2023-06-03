package product

import (
	"server/global"
	"time"

	"gorm.io/gorm"
)

// 产品表
/*
Create table products
(
	id bigint primary key auto_increment,
	name varchar(20) not null unique,
	supplier_phone varchar(20) not null,
	create_time timestamp with time zone not null,
	foreign key (supplier_name) references suppliers(name) on update cascade on delete restrict,
);
*/
type Product struct {
	ID           int       `json:"id" gorm:"primary_key"`
	Name         string    `json:"name" gorm:"type:varchar(20);unique;not null"`
	SupplierName string    `json:"supplier_name" gorm:"type:varchar(20);not null"`
	Supplier     Supplier  `json:"supplier" gorm:"foreignkey:SupplierName;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreateTime   time.Time `json:"create_time" gorm:"not null"`
	UpdateTime   time.Time `json:"update_time" gorm:"not null"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	p.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	p.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	p.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

// 供应商表
/*
Create table suppliers
(
	name varchar(20) primary key,
	phone varchar(20) not null unique,
	locations varchar(20) not null
)
*/
type Supplier struct {
	Name     string `json:"name" gorm:"type:varchar(20);primary_key"`
	Phone    string `json:"phone" gorm:"type:char(11);unique;check:,(phone ~ '^[0-9]{11}$');"`
	Location string `json:"location" gorm:"type:varchar(20);not null"`
}

func InitializeProduct(db *gorm.DB) {
	// 录入供应商信息
	/*
		INSERT INTO suppliers (name, phone, locations) VALUES
		('Apple', '14337281938', 'california')
	*/
	Suppliers := []Supplier{
		{Name: "Apple", Phone: "14337281938", Location: "california"},
		{Name: "Huawei", Phone: "14412343093", Location: "Shenzhen"},
		{Name: "Samsung", Phone: "14262902578", Location: "seoul"},
		{Name: "Microsoft", Phone: "17890124386", Location: "Washington"},
		{Name: "Google", Phone: "13523403598", Location: "California"},
		{Name: "Sony", Phone: "13452306238", Location: "Tokyo"},
		{Name: "Amazon", Phone: "19423103450", Location: "Washington"},
		{Name: "nintendo", Phone: "13402382311", Location: "Kyoto"},
		{Name: "Xiaomi", Phone: "13209327391", Location: "Beijing"},
		{Name: "OPPO", Phone: "13125033292", Location: "Beijing"},
		{Name: "HuaqiangNorth", Phone: "15242091839", Location: "Shenzhen"},
	}
	global.GL_DB.Create(&Suppliers)

	// 录入产品信息
	/*
		INSERT INTO products (name, supplier_name, create_time) VALUES
		('iPhone 13', 'Apple', '2021-09-14 00:00:00'),
	*/
	Products := []Product{
		{Name: "iPhone 13", SupplierName: "Apple"},
		{Name: "MacBook Pro", SupplierName: "Apple"},
		{Name: "Mate 50", SupplierName: "Huawei"},
		{Name: "HuaWei Watch 3", SupplierName: "Huawei"},
		{Name: "HuaWei MatePad Pro", SupplierName: "Huawei"},
		{Name: "Galaxy S22", SupplierName: "Samsung"},
		{Name: "Surface Pro 8", SupplierName: "Microsoft"},
		{Name: "Pixel 6", SupplierName: "Google"},
		{Name: "PlayStation 5", SupplierName: "Sony"},
		{Name: "Kindle", SupplierName: "Amazon"},
		{Name: "Switch", SupplierName: "nintendo"},
		{Name: "Xiaomi 12", SupplierName: "Xiaomi"},
		{Name: "Xiaomi Pad 5", SupplierName: "Xiaomi"},
		{Name: "Xiaomi Watch", SupplierName: "Xiaomi"},
		{Name: "OPPO Find X5", SupplierName: "OPPO"},
		{Name: "OPPO Watch 2", SupplierName: "OPPO"},
	}
	global.GL_DB.Create(&Products)
}
