package main

import (
	"fmt"
	"fynegui/lib"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// appsTable Data
var appsList = [][]string{
	{"ID", "Name", "Email"},
	{"1", "John Doe", "john@example.com"},
	{"2", "Jane Smith", "jane@example.com"},
	{"3", "Emily Johnson", "emily@example.com"},
	{"4", "Michael Brown", "michael@example.com"},
	{"5", "Jessica Williams", "jessica@example.com"},
}

var AppInfos = []lib.App{
	lib.AppInfo{Id: 1, Name: "WebStrom", Source: "dir/cmd", Type: "APP"},
	lib.AppInfo{Id: 2, Name: "WebStrom", Source: "dir/cmd", Type: "APP"},
	lib.AppInfo{Id: 3, Name: "WebStrom", Source: "dir/cmd", Type: "APP"},
}

var apps = []lib.App{
	lib.AppInfo{Id: 1, Name: "App One", Source: "Source A", Type: "Type X"},
	lib.AppInfo{Id: 2, Name: "App Two", Source: "Source B", Type: "Type Y"},
	lib.AppInfo{Id: 3, Name: "App Three", Source: "Source C", Type: "Type Z"},
	lib.AppInfo{Id: 4, Name: "App Four", Source: "Source D", Type: "Type X"},
	lib.AppInfo{Id: 5, Name: "App Five", Source: "Source E", Type: "Type Y"},
}

func MakeLabel(name string) *widget.Label {
	return widget.NewLabel(name)
}

func MakeButton(name string) *widget.Button {
	return widget.NewButton(name, func() {
		log.Println("Clicked", name)
	})
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("App Starter")
	myWindow.Resize(fyne.NewSize(600, 400))

	label := MakeLabel("Added Apps")
	label.TextStyle = fyne.TextStyle{Bold: true}

	addButton := MakeButton("Add Apps")

	hBox := container.NewHBox(label, layout.NewSpacer(), addButton)

	appsTable := widget.NewTable(
		func() (int, int) {
			return len(apps) + 1, 5 // +1 for header row, 5 columns (Id, Name, Source, Type, Actions)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("wide content"))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 { // Header row
				switch i.Col {
				case 0:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("ID")}
				case 1:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Name")}
				case 2:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Source")}
				case 3:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Type")}
				case 4:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Actions")}
				}
			} else { // Data rows
				app := apps[i.Row-1] // Adjust index because of header row
				switch i.Col {
				case 0:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(app.GetDetail())}
				case 1:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(app.GetDetail())}
				case 2:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(app.GetDetail())}
				case 3:
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(app.GetDetail())}
				case 4:
					startIcon := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
						fmt.Printf("Start action for row %d\n", i.Row)
					})
					closeIcon := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
						fmt.Printf("Close action for row %d\n", i.Row)
					})
					o.(*fyne.Container).Objects = []fyne.CanvasObject{startIcon, closeIcon}
				}
			}
		})
	/* // appsTable with Action Header
	appsTable := widget.NewTable(
		func() (int, int) {
			return len(appsList), len(appsList[0]) + 1 // +1 for action column
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("wide content"))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(appsList[0]) { // Action column
				if i.Row == 0 {
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Actions")}
				} else {
					startIcon := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
						fmt.Printf("Start action for row %d\n", i.Row)
					})
					closeIcon := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
						fmt.Printf("Close action for row %d\n", i.Row)
					})
					o.(*fyne.Container).Objects = []fyne.CanvasObject{startIcon, closeIcon}
				}
			} else {
				o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(appsList[i.Row][i.Col])}
			}
			o.Refresh() // Refresh the container to apply changes
		}) */

	// Ensure the appsTable expands to fill the available space
	appsTable.SetColumnWidth(0, 50)  // Set a fixed width for the first column
	appsTable.SetColumnWidth(1, 150) // Set a fixed width for the second column
	appsTable.SetColumnWidth(2, 250) // Set a fixed width for the third column
	appsTable.SetColumnWidth(3, 100) // Set a fixed width for the action column
	appsTable.SetColumnWidth(3, 100) // Set a fixed width for the action column

	// Scrollable appsTable with Fixed Height
	scrollContainer := container.NewVScroll(appsTable)
	scrollContainer.SetMinSize(fyne.NewSize(600, 200)) // Set fixed height for the scrollable appsTable

	// Footer with "Start All" Button
	startAllButton := MakeButton("Start All")
	footer := container.NewHBox(layout.NewSpacer(), startAllButton)

	// Main Container
	mainContainer := container.NewBorder(hBox, footer, nil, nil, scrollContainer)
	myWindow.SetContent(mainContainer)

	myWindow.ShowAndRun()
}
