package dialogs

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
)

// Configuration window.
func ConfigWindow(win *fyne.Window, app fyne.App, refresh func()) {
	var selGridStyle, selGridSize string

	// selector to determine grid size
	sizeSelector := widget.NewRadioGroup([]string{
		globals.SIZE_LARGE,
		globals.SIZE_DEFAULT,
		globals.SIZE_SMALL,
	}, func(sel string) {
		selGridSize = sel
	})
	sizeSelector.SetSelected(globals.GridSize)

	// entry to display and set the configured paths
	input := newEntryPaths()

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
				paths := formatInput(input.Text)
				f := utils.SetConfig(paths)
				f.Close()

				// refresh the main window
				refresh()
			}
		}, *win)
	dia.Resize(fyne.NewSize(400, 400))

	dia.Show()
}

// format the input of the user.
func formatInput(s string) string {
	var res string

	for _, i := range strings.Split(s, "\n") {
		aux := strings.TrimSpace(i)
		strings.Replace(aux, ",", "", 1)
		res += aux
	}

	return res
}

// Creates a new input field to display the configured paths inside.
func newEntryPaths() *widget.Entry {
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
