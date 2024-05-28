package main

import (
	"fynegui/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	appInstance := app.New()
	ui.HomeWindow(appInstance).Show()
	appInstance.Run()
}
