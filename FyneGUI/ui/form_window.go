package ui

import (
	"fmt"
	"fynegui/lib"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowFormWindow(windowRef fyne.Window) {

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter script name")

	sourceEntry := widget.NewEntry()
	sourceEntry.SetPlaceHolder("Selct file path")
	sourceEntry.Disable()

	sourceButton := widget.NewButton("Browse", func() {
		fileDialog := dialog.NewFolderOpen(func(file fyne.ListableURI, err error) {
			if file == nil {
				return
			} else {
				sourceEntry.SetText(string(file.Path()))
			}
		}, windowRef)

		fileDialog.Show()

	})

	form := container.NewVBox(
		container.NewVBox(widget.NewLabel("Name:"), nameEntry),
		layout.NewSpacer(),
		container.NewVBox(widget.NewLabel("Select Path:"), sourceEntry),
		container.NewVBox(sourceButton),
	)

	dialog.ShowCustomConfirm("Submit Script", "Submit", "Cancel", form, func(b bool) {
		if b {
			script := lib.Script{
				Name:   nameEntry.Text,
				Source: sourceEntry.Text,
			}
			fmt.Println(script)
		}

	}, windowRef)
}
