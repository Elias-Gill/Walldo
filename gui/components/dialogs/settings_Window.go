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
		pathInput.SetText("")

		pathInput.Refresh()
		pathsList.Refresh()
	}
	pathInput.Resize(fyne.NewSize(200, 500))

	// open the file explorer to select a folder
	pathPickerButton := widget.NewButton("Open explorer", func() {
		NewPathPicker(*win, func(path string) {
			data = append(data, path)
			pathInput.SetText("")
			pathInput.Refresh()
			pathsList.Refresh()
		})
	})

	pathsList.Resize(fyne.NewSize(float32(pathsList.Size().Width), 25))

	// Window content
	layout := container.NewGridWithRows(2,
		container.NewVBox(sizeSelector, widget.NewSeparator()),
		container.NewGridWithRows(2, pathsList, container.NewVBox(pathPickerButton)),
	)

	// Create the new dialog window (the main container)
	dia := dialog.NewCustomConfirm("Settings", "Confirm", "Cancel", layout,
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
