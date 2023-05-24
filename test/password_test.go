package test

import (
	"server/util"
	"testing"
)

func TestPassword(t *testing.T) {
	password := "123456"
	hashPassword, err := util.HashAndSalt(password)
	if err != nil {
		t.Error(err)
	}
	isValid := util.CheckPasswordHash(password, hashPassword)
	if !isValid {
		t.Error("Password is not valid")
	}
}
