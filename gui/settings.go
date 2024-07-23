package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
)

const (
	large  = "Large"
	normal = "Default"
	small  = "Small"
)

var names = map[config.GridSize]string{
	config.LARGE:  large,
	config.NORMAL: normal,
	config.SMALL:  small,
}

var sizes = map[string]config.GridSize{
	large:  config.LARGE,
	normal: config.NORMAL,
	small:  config.SMALL,
}

// Configuration window.
//
//nolint:all
func newConfigWindow(refresh func()) {
	var selGridSize string

	// grid size selector
	sizeSelector := widget.NewRadioGroup([]string{
		small, normal, large,
	}, func(sel string) {
		selGridSize = sel
	})
	sizeSelector.SetSelected(names[config.GetGridSize()])

	// path list
	data := config.GetPaths()
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
		NewPathPicker(config.GetWindow(), func(path string) {
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
				config.SetGridSize(sizes[selGridSize])
				config.SetPaths(data)

				// refresh the main window
				refresh()
			}
		}, config.GetWindow())

	dia.Resize(fyne.NewSize(
		config.GetWindow().Canvas().Size().Width,
		config.GetWindow().Canvas().Size().Height),
	)

	dia.Show()
}
