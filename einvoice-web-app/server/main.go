package main

import "github.com/slovak-egov/einvoice/einvoice-web-app/server/app"

func main() {
	a := app.App{}

	a.Initialize()

	a.Run()
}
