package main

import (
	"fynegui/lib"
	"fynegui/ui"

	"fyne.io/fyne/v2/app"
)

var apps = []lib.App{
	lib.AppInfo{Id: 1, Name: "App One", Source: "Source A"},
	lib.AppInfo{Id: 2, Name: "App Two", Source: "Source B"},
	lib.AppInfo{Id: 3, Name: "App Three", Source: "Source C"},
	lib.Script{Id: 4, Name: "Script", Source: "./hello"},
	lib.Script{Id: 5, Name: "My Works", Source: "./works"},
}

func main() {
	appInstance := app.New()
	ui.HomeWindow(appInstance, apps).ShowAndRun()
}
