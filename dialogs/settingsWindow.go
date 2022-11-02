package dialogs

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
)

// Configuration window
func ConfigWindow(win *fyne.Window, app fyne.App, refresh *widget.Button) {
	var selGridStyle string
	var selGridSize string

	// selector to determine grid size
	grid_size_selector := widget.NewRadioGroup([]string{"large", "default", "small"},
		func(sel string) {
			selGridSize = sel
		})
	grid_size_selector.SetSelected(globals.GridSize)

	// grid style selector (Borderless or captions)
	gridStyleSelector := widget.NewRadioGroup(
		[]string{"Borderless", "Captions"},
		func(sel string) {
			selGridStyle = sel
		})
	gridStyleSelector.SetSelected(globals.GridTitles)

	// entry to display and set the configured paths
	input := newEntryPaths()

	// Window content
	cont := []*widget.FormItem{
		widget.NewFormItem("Images size", grid_size_selector),
		widget.NewFormItem("Grid style", gridStyleSelector),
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
				globals.GridTitles = selGridStyle
				globals.GridSize = selGridSize

				// update configured paths
				paths := formatInput(input.Text)
				f := utils.SetConfig(paths)
				f.Close()

				// refresh the main window
				refresh.OnTapped()
			}
		}, *win)
	dia.Resize(fyne.NewSize(400, 400))

	dia.Show()
}

// format the input of the user
func formatInput(s string) string {
	var res string
	for _, i := range strings.Split(s, "\n") {
		aux := strings.TrimSpace(i)
		strings.Replace(aux, ",", "", 1)
		res += aux
	}
	return res
}

// function to create a new entry for changing the configured paths
func newEntryPaths() *widget.Entry {
	input := widget.NewEntry()
	input.MultiLine = true
	input.SetPlaceHolder(`C:/User/fondos, \nC:/Example/images`)

	// get the current configured paths and display them
	var c string
	aux := utils.GetConfiguredPaths()
	for count, i := range aux {
		c += i
		if count < len(aux)-1 {
			c += ",\n"
		}
	}

	if c != "" {
		input.SetText(c)
	}

	return input
}
