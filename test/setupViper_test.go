package test

import (
	"testing"

	"server/global"
	"server/setup"
)

func TestInitializeViper(t *testing.T) {
	global.GL_VIPER = setup.InitializeViper(".././")
	t.Logf("%#v\n", global.GL_CONFIG)
}
