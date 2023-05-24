package test

import (
	"testing"

	"server/setup"
)

func TestInitializeDB(t *testing.T) {
	setup.InitializeViper(".././")
	setup.InitializeDB()
}
