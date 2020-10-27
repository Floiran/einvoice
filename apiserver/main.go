package main

import (
	"github.com/slovak-egov/einvoice/apiserver/app"
	"github.com/slovak-egov/einvoice/apiserver/config"
)

func main() {
	appConfig := config.Init()

	a := app.App{Config: appConfig}

	a.Initialize()

	a.Run()

	a.Close()
}
