package test

import (
	"fmt"
	"os/exec"
	"testing"

	"server/global"
	"server/setup"
)

func TestInitializeViper(t *testing.T) {
	global.GL_VIPER = setup.InitializeViper(".././")
}

func TestCmd(t *testing.T) {
	cmd := exec.Command("make", "-C", "../", "fmt")

	out, err := cmd.Output()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(out))
}
