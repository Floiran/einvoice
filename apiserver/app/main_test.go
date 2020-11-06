package app_test

import (
	"os"
	"testing"

	"github.com/slovak-egov/einvoice/apiserver/app"
)

var a *app.App

func TestMain(m *testing.M) {
	a = &app.App{}

	a.Initialize()

	result := m.Run()

	a.Close()

	os.Exit(result)
}
