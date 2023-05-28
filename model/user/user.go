package user

import (
	"time"

	"gorm.io/gorm"

	"server/global"
	"server/util"
)

// 用户表
/*
Create table user
(
	username varchar(20) primary key,
	password char(60) not null,
	role varchar(10) not null,
	create_time timestamp with time zone not null,
	update_time timestamp with time zone not null
);
*/
type User struct {
	Username string `json:"username" gorm:"type:varchar(20);primary_key;"`
	Password string `json:"password" gorm:"type:char(60);not null"` // bcrypt hashed
	Role     string `json:"role" gorm:"type:varchar(10);not null"`

	CreateTime time.Time `json:"create_time" gorm:"not null"`
	UpdateTime time.Time `json:"update_time" gorm:"not null"`
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

func InitializeUser(db *gorm.DB) {
	u1 := User{
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	}

	u2 := User{
		Username: "xypf",
		Password: "xypf",
		Role:     "editor",
	}

	u3 := User{
		Username: "fzjc",
		Password: "fzjc",
		Role:     "analyst",
	}

	u4 := User{
		Username: "guest",
		Password: "guest",
		Role:     "guest",
	}

	// 加密
	u1Pw, _ := util.HashAndSalt(u1.Password)
	u2Pw, _ := util.HashAndSalt(u2.Password)
	u3Pw, _ := util.HashAndSalt(u3.Password)
	u4Pw, _ := util.HashAndSalt(u4.Password)

	u1.Password = u1Pw
	u2.Password = u2Pw
	u3.Password = u3Pw
	u4.Password = u4Pw

	/*
		INSERT INTO users (username,password,role,create_time,update_time)
		VALUES ('admin', 'admin', 'admin', '2021-01-01 00:00:00', '2021-01-01 00:00:00')
	*/

	users := []User{u1, u2, u3, u4}

	global.GL_DB.Model(&User{}).Create(&users)
}
