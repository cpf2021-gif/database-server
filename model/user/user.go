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
	sex char(1) not null check (sex in ('女', '男')),
	phone varchar(11) not null check (phone ~ '^[0-9]{11}$') unique,
	role varchar(10) not null,
	create_time timestamp with time zone not null,
	update_time timestamp with time zone not null
);
*/
type User struct {
	Username string `json:"username" gorm:"type:varchar(20);primary_key;"`
	Password string `json:"password" gorm:"type:char(60);not null"` // bcrypt hashed
	Sex      string `json:"sex" gorm:"type:char(1);not null;check:,sex in ('女', '男')"`
	Phone    string `json:"phone" gorm:"type:char(11);not null;check:,(phone ~ '^[0-9]{11}$');unique"`
	Role     string `json:"role" gorm:"type:varchar(10);not null;check:,role in ('admin', 'editor', 'analyst', 'guest')"`

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
		Sex:      "男",
		Phone:    "12345678901",
	}

	u2 := User{
		Username: "xypf",
		Password: "xypf",
		Role:     "editor",
		Sex:      "男",
		Phone:    "12345678902",
	}

	u3 := User{
		Username: "fzjc",
		Password: "fzjc",
		Role:     "analyst",
		Sex:      "女",
		Phone:    "12345678903",
	}

	u4 := User{
		Username: "guest",
		Password: "guest",
		Sex:      "女",
		Role:     "guest",
		Phone:    "12345678904",
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
		INSERT INTO users (username,password,role,sex,phone,create_time,update_time)
		VALUES ('admin', 'admin', 'admin', '男', '12345678901', '2021-01-01 00:00:00', '2021-01-01 00:00:00')
	*/

	users := []User{u1, u2, u3, u4}

	global.GL_DB.Model(&User{}).Create(&users)
}
