package supplier

// 供应商表
type Supplier struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"type:varchar(20);not null"`
	Phone string `json:"phone" gorm:"type:varchar(20);not null"`
}
