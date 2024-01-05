package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
)

// Configuration window.
//
//nolint:all
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

	// path list
	data := utils.GetConfiguredPaths()
	pathsList := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	pathsList.OnSelected = func(id int) {
		// delete the selected element
		var aux []string
		for i := 0; i < len(data); i++ {
			if i != id {
				aux = append(aux, data[i])
			}
		}
		data = aux
		pathsList.Refresh()
	}

	// path pathInput
	pathInput := widget.NewEntry()
	pathInput.MultiLine = false
	pathInput.SetPlaceHolder(`C:/User/user/fondos`)
	pathInput.OnSubmitted = func(t string) {
		data = append(data, t)
		pathInput.Text = ""

		pathInput.Refresh()
		pathsList.Refresh()
	}

	// Window content
	cont := []*widget.FormItem{
		widget.NewFormItem("Images size", sizeSelector),
		widget.NewFormItem("", pathInput),
		widget.NewFormItem("", container.NewMax(pathsList)),
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
				utils.MustWriteConfig(utils.
					NewConfig().
					WithPaths(data),
				)

				// refresh the main window
				refresh()
			}
		}, *win)

	dia.Resize(fyne.NewSize(
		globals.Window.Canvas().Size().Width,
		globals.Window.Canvas().Size().Height),
	)

	dia.Show()
}
