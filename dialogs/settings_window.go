package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
)

// Configuration window
func ConfigWindow(win *fyne.Window, app fyne.App, refresh *widget.Button) {
	var sel_grid_style string
	var sel_grid_size string
	var sel_layout_style string

	// selector to determine grid size
	grid_size_selector := widget.NewRadioGroup([]string{"large", "default", "small"},
		func(sel string) {
			sel_grid_size = sel
		})
	grid_size_selector.SetSelected(globals.GridSize)

	// grid style selector (Borderless or captions)
	grid_style_selector := widget.NewRadioGroup(
		[]string{"Borderless", "Captions"},
		func(sel string) {
			sel_grid_style = sel
		})
	grid_style_selector.SetSelected(globals.GridTitles)

	// Layout style selector (grid or rows)
	layout_selector := widget.NewRadioGroup(
		[]string{"Grid", "Rows"},
		func(sel string) {
			sel_layout_style = sel
		})
	layout_selector.SetSelected(globals.LayoutStyle)

    // entry to display and set the configured paths
    input := newEntryPaths()

	// Window content
	cont := []*widget.FormItem{
		widget.NewFormItem("Layout", layout_selector),
		widget.NewFormItem("Images size", grid_size_selector),
		widget.NewFormItem("Grid style", grid_style_selector),
		widget.NewFormItem("", input),
	}

	// Create the new dialog window (the main container)
	dia := dialog.NewForm("Settings", "Confirm", "Cancel", cont,
        // function to refresh the content with the new given config
		func(status bool) {
			if status {
				// update fyne config API
				globals.MyApp.Preferences().SetString("GridTitles", sel_grid_style)
				globals.MyApp.Preferences().SetString("GridSize", sel_grid_size)
				globals.MyApp.Preferences().SetString("Layout", sel_layout_style)

				// update global variables
				globals.GridTitles = sel_grid_style
				globals.GridSize = sel_grid_size
				globals.LayoutStyle = sel_layout_style

                // update configured paths
                f := utils.SetConfig(input.Text)
                f.Close()

				// refresh the main window
				refresh.OnTapped()
			}
		}, *win)
    dia.Resize(fyne.NewSize(400, 400))

	dia.Show()
}

// function to create a new entry for changing the configured paths
func newEntryPaths () *widget.Entry{
    input := widget.NewEntry()
    input.SetPlaceHolder(`C:/User/fondos, C:/Example/images`)

    // get the current configured paths and display them
    var c string
    aux := utils.GetConfiguredPaths()
    for count, i := range aux{
        c += i 
        if (count < len(aux)-1){
            c += ", "
        }
    }

    if c != ""{
        input.SetText(c)
    }

    return input
}
