package user

import (
	"time"

	"gorm.io/gorm"
)

// 用户表
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"type:varchar(20);not null;unique"`
	Password string `json:"password" gorm:"type:char(60);not null"` // bcrypt hashed
	Role     string `json:"role" gorm:"type:varchar(10);not null"`

	CreateTime time.Time `json:"create_time" gorm:"index"`
	UpdateTime time.Time `json:"update_time" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	u.CreateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	u.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	u.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	return nil
}
