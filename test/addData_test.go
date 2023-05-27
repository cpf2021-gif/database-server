package test

import (
	"testing"

	"server/global"
	"server/model/user"
	"server/setup"
	"server/util"
)

func TestUser(t *testing.T) {
	setup.InitializeViper(".././")
	global.GL_DB = setup.InitializeDB()

	u := user.User{
		Username: "xypf",
		Password: "xypf",
		Role:     "editor",
	}

	hashedPassword, err := util.HashAndSalt(u.Password)
	if err != nil {
		t.Error(err)
	}

	u.Password = hashedPassword
	if global.GL_DB.Model(&user.User{}).Create(&u).Error != nil {
		t.Error("create user error")
	}
}
