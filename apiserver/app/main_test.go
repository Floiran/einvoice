package app

import (
	"os"
	"testing"

	"github.com/slovak-egov/einvoice/apiserver/config"
)

var a App

func TestMain(m *testing.M) {
	testConfig := config.Init()

	a = App{Config: testConfig}

	a.Initialize()

	result := m.Run()

	a.Close()

	os.Exit(result)
}
