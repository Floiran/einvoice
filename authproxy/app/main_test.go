package app

import (
	"os"
	"testing"

	"github.com/slovak-egov/einvoice/authproxy/config"
	"github.com/slovak-egov/einvoice/authproxy/manager"
)

var a *App

func TestMain(m *testing.M) {
	a = &App{}

	a.config = config.Init()
	a.manager = manager.Init(a.config)
	// TODO: initialize mocked SlovenskoSk
	a.InitializeRouter()

	result := m.Run()

	a.Close()

	os.Exit(result)
}
