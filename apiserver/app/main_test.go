package app

import (
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}

	a.Initialize()

	result := m.Run()

	a.Close()

	os.Exit(result)
}
