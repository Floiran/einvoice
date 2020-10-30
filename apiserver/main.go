package main

import "github.com/slovak-egov/einvoice/apiserver/app"

func main() {
	a := app.App{}

	a.Initialize()

	defer a.Close()

	a.Run()
}
