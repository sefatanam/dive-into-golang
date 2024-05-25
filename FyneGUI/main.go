package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeLabel(name string) *widget.Label {
	return widget.NewLabel(name)
}

func MakeButton(name string) *widget.Button {
	return widget.NewButton(name, func() {
		log.Println("Clicked", name)
	})
}

func main() {
	/* myApp := app.New()
	myWindow := myApp.NewWindow("App Starter")
	myWindow.Resize(fyne.NewSize(600, 400))

	label := widget.NewLabel("Added Applications")
	addButton := widget.NewButton("Add Apps", func() { fmt.Println("test") })

	customers := [][]string{
		{"ID", "Name", "Email"},
		{"1", "John Doe", "john@example.com"},
		{"2", "Jane Smith", "jane@example.com"},
		{"3", "Emily Johnson", "emily@example.com"},
	}

	list := widget.NewTable(
		func() (int, int) {
			return 4, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {

			o.(*widget.Label).SetText(customers[i.Row][i.Col])
		})

	labelButtonContainer := container.New(layout.NewHBoxLayout(), label, layout.NewSpacer(), addButton)
	tableContainer := container.New(layout.NewHBoxLayout(), container.NewVScroll(list))

	myWindow.SetContent(container.NewVBox(labelButtonContainer, tableContainer))
	myWindow.ShowAndRun() */

	/* 	myApp := app.New()
	   	myWindow := myApp.NewWindow("App Starter")
	   	myWindow.Resize(fyne.NewSize(600, 400))

	   	// Header with Label and Button
	   	label := widget.NewLabel("Added Applications")
	   	addButton := widget.NewButton("Add Apps", func() {
	   		fmt.Println("test")
	   	})

	   	hBox := container.NewHBox(label, layout.NewSpacer(), addButton) */

	/* // List of Applications
	apps := []string{"1. WebStorm", "2. Rider", "1. WebStorm", "1. WebStorm", "1. WebStorm"}

	list := widget.NewList(
		func() int {
			return len(apps)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(apps[i])
		})

	// Scrollable List with Fixed Height
	scrollContainer := container.NewVScroll(list) */

	// Replace this part for List of Applications with a Table of Applications
	/* customers := [][]string{
		{"ID", "Name", "Email", "Actions"},
		{"1", "John Doe", "john@example.com", "Start"},
		{"2", "Jane Smith", "jane@example.com", "Start"},
		{"3", "Emily Johnson", "emily@example.com", "Start"},
		{"4", "Michael Brown", "michael@example.com", "Start"},
		{"5", "Jessica Williams", "jessica@example.com", "Start"},
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(customers), len(customers[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(customers[i.Row][i.Col])
		})

	// Scrollable Table with Fixed Height
	scrollContainer := container.NewVScroll(table)
	scrollContainer.SetMinSize(fyne.NewSize(600, 200)) // Set fixed height for the scrollable table

	scrollContainer.SetMinSize(fyne.NewSize(600, 200)) // Set fixed height for the scrollable list

	// Footer with "Start All" Button
	startAllButton := widget.NewButton("Start All", func() {
		fmt.Println("Starting all applications...")
	})

	footer := container.NewHBox(layout.NewSpacer(), startAllButton)

	// Main Container
	mainContainer := container.NewVBox(hBox, scrollContainer, layout.NewSpacer(), footer)
	myWindow.SetContent(mainContainer)

	myWindow.ShowAndRun() */

	/* myApp := app.New()
	myWindow := myApp.NewWindow("App Starter")
	myWindow.Resize(fyne.NewSize(600, 400))

	// Header with Label and Button
	label := widget.NewLabel("Added Applications")
	addButton := widget.NewButton("Add Apps", func() {
		fmt.Println("Add Apps button clicked")
	})

	hBox := container.NewHBox(label, layout.NewSpacer(), addButton)

	// Table Data
	customers := [][]string{
		{"ID", "Name", "Email"},
		{"1", "John Doe", "john@example.com"},
		{"2", "Jane Smith", "jane@example.com"},
		{"3", "Emily Johnson", "emily@example.com"},
		{"4", "Michael Brown", "michael@example.com"},
		{"5", "Jessica Williams", "jessica@example.com"},
	}

	// Table with Action Header
	table := widget.NewTable(
		func() (int, int) {
			return len(customers), len(customers[0]) + 1 // +1 for action column
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("wide content"))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(customers[0]) { // Action column
				if i.Row == 0 {
					o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel("Actions")}
				} else {
					startIcon := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
						fmt.Printf("Start action for row %d\n", i.Row)
					})
					closeIcon := widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {
						fmt.Printf("Close action for row %d\n", i.Row)
					})
					o.(*fyne.Container).Objects = []fyne.CanvasObject{startIcon, closeIcon}
				}
			} else {
				o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(customers[i.Row][i.Col])}
			}
		})

	// Scrollable Table with Fixed Height
	scrollContainer := container.NewVScroll(table)
	scrollContainer.SetMinSize(fyne.NewSize(600, 200)) // Set fixed height for the scrollable table

	// Footer with "Start All" Button
	startAllButton := widget.NewButton("Start All", func() {
		fmt.Println("Starting all applications...")
	})

	footer := container.NewHBox(layout.NewSpacer(), startAllButton)

	// Main Container
	mainContainer := container.NewVBox(hBox, scrollContainer, layout.NewSpacer(), footer)
	myWindow.SetContent(mainContainer)

	myWindow.ShowAndRun() */

	myApp := app.New()
	myWindow := myApp.NewWindow("App Starter")
	myWindow.Resize(fyne.NewSize(600, 400))

	// Header with Label and Button
	label := widget.NewLabel("Added Applications")
	addButton := widget.NewButton("Add Apps", func() {
		fmt.Println("Add Apps button clicked")
	})

	hBox := container.NewHBox(label, layout.NewSpacer(), addButton)

	// Table Data
	customers := [][]string{
		{"ID", "Name", "Email"},
		{"1", "John Doe", "john@example.com"},
		{"2", "Jane Smith", "jane@example.com"},
		{"3", "Emily Johnson", "emily@example.com"},
		{"4", "Michael Brown", "michael@example.com"},
		{"5", "Jessica Williams", "jessica@example.com"},
	}

	// Table with Action Header
	table := widget.NewTable(
		func() (int, int) {
			return len(customers), len(customers[0]) + 1 // +1 for action column
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("wide content"))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(customers[0]) { // Action column
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
				o.(*fyne.Container).Objects = []fyne.CanvasObject{widget.NewLabel(customers[i.Row][i.Col])}
			}
			o.Refresh() // Refresh the container to apply changes
		})

	// Ensure the table expands to fill the available space
	table.SetColumnWidth(0, 50)  // Set a fixed width for the first column
	table.SetColumnWidth(1, 150) // Set a fixed width for the second column
	table.SetColumnWidth(2, 250) // Set a fixed width for the third column
	table.SetColumnWidth(3, 150) // Set a fixed width for the action column

	// Scrollable Table with Fixed Height
	scrollContainer := container.NewVScroll(table)
	scrollContainer.SetMinSize(fyne.NewSize(600, 200)) // Set fixed height for the scrollable table

	// Footer with "Start All" Button
	startAllButton := widget.NewButton("Start All", func() {
		fmt.Println("Starting all applications...")
	})

	footer := container.NewHBox(layout.NewSpacer(), startAllButton)

	// Main Container
	mainContainer := container.NewBorder(hBox, footer, nil, nil, scrollContainer)
	myWindow.SetContent(mainContainer)

	myWindow.ShowAndRun()
}
