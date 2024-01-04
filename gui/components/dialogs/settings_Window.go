package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
)

// Configuration window.
func ConfigWindow(win *fyne.Window, app fyne.App, refresh func()) {
	var selGridStyle, selGridSize string

	// grid size selector
	sizeSelector := widget.NewRadioGroup([]string{
		globals.SIZE_LARGE,
		globals.SIZE_DEFAULT,
		globals.SIZE_SMALL,
	}, func(sel string) {
		selGridSize = sel
	})
	sizeSelector.SetSelected(globals.GridSize)

	// entry to display and set the configured paths
	input := newPathsInput()

	// Window content
	cont := []*widget.FormItem{
		widget.NewFormItem("Images size", sizeSelector),
		widget.NewFormItem("", input),
	}

	// Create the new dialog window (the main container)
	dia := dialog.NewForm("Settings", "Confirm", "Cancel", cont,
		// function to refresh the content with the new given config
		func(status bool) {
			if status {
				// update fyne config API
				globals.MyApp.Preferences().SetString("GridTitles", selGridStyle)
				globals.MyApp.Preferences().SetString("GridSize", selGridSize)

				// update global variables
				globals.GridSize = selGridSize

				// update configured paths
				config := utils.NewConfig()
				config.WithPaths(formatInput(input.Text))
				utils.MustWriteConfig(config)

				// refresh the main window
				refresh()
			}
		}, *win)
	dia.Resize(fyne.NewSize(400, 400))

	dia.Show()
}

// Creates a new input field to display the configured paths inside.
func newPathsInput() *widget.Entry {
	input := widget.NewEntry()
	input.MultiLine = true
	input.SetPlaceHolder(`C:/User/fondos, \nC:/Example/images`)

	// format the configured paths for the fyne input
	paths := ""

	p := utils.GetConfiguredPaths()
	for i, path := range p {
		paths += path
		if i < len(p)-1 { // to avoid putting a extra line jump at the end
			paths += ",\n"
		}
	}

	input.SetText(paths)

	return input
}

func NewPathsList() fyne.Widget {
	data := utils.GetConfiguredPaths()

	resultsWidget := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

    // TODO: delete function
	resultsWidget.OnSelected = func(id int) {}

	return resultsWidget
}
