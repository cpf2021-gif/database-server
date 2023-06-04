package test

import (
	"fmt"
	"testing"

	"server/global"
	"server/model/user"
	"server/setup"
	// "server/util"
)

func TestUser(t *testing.T) {
	setup.InitializeViper(".././")
	global.GL_DB = setup.InitializeDB()

	// u := user.User{
	// 	Username: "xypf",
	// 	Password: "xypf",
	// 	Role:     "editor",
	// }

	// hashedPassword, err := util.HashAndSalt(u.Password)
	// if err != nil {
	// 	t.Error(err)
	// }

	// u.Password = hashedPassword
	// if global.GL_DB.Model(&user.User{}).Create(&u).Error != nil {
	// 	t.Error("create user error")
	// }
}

func TestGetUsers(t *testing.T) {
	global.GL_VIPER = setup.InitializeViper(".././")
	global.GL_DB = setup.InitializeDB()

	var users []user.User

	if global.GL_DB.Model(&user.User{}).Find(&users).Error != nil {
		t.Errorf("get users error")
	}

	for i := range users {
		users[i].CreateTime = users[i].CreateTime.UTC()
		users[i].UpdateTime = users[i].UpdateTime.UTC()
	}

	fmt.Println(users)
}
