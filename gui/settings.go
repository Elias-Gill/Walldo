package gui

import (
	"os"
	"runtime"
	"strings"

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

// Generates and displays the configuration window.
//
// The callback function that gets executed after the configuration
// has been successfully updated, is intended to refresh the wallpaper gallery and update
// other UI elements in response to configuration changes.
func newConfigWindow(callback func()) {
	var selectedGridSize string

	// grid size selector
	sizeSelector := widget.NewRadioGroup([]string{
		small, normal, large,
	}, func(sel string) {
		selectedGridSize = sel
	})
	sizeSelector.SetSelected(names[config.GetGridSize()])
	selectorLabel := widget.NewLabel("Image grid cell size")
	selectorLabel.TextStyle.Bold = true
	sizeSelectorContainer := container.NewVBox(selectorLabel, sizeSelector)

	// path list
	listedPaths := config.GetRawSearchPaths()
	pathsList := widget.NewList(
		func() int {
			return len(listedPaths)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(listedPaths[i])
		})

	pathsList.OnSelected = func(id int) {
		// delete the selected element
		var aux []string
		for i := 0; i < len(listedPaths); i++ {
			if i != id {
				aux = append(aux, listedPaths[i])
			}
		}
		listedPaths = aux
		pathsList.Refresh()
	}

	// Wallpapers paths selection
	pathsSelectorLabel := widget.NewLabel("Wallpapers paths")
	pathsSelectorLabel.TextStyle.Bold = true

	// input
	pathInput := widget.NewEntry()
	pathInput.MultiLine = false

	if runtime.GOOS == "windows" {
		pathInput.SetPlaceHolder(`C:/Users/user/wallpapers`)
	} else {
		pathInput.SetPlaceHolder(`~/wallpapers`)
	}

	pathInput.OnSubmitted = func(t string) {
		listedPaths = append(listedPaths, t)
		pathInput.SetText("")

		pathInput.Refresh()
		pathsList.Refresh()
	}
	pathInput.Resize(fyne.NewSize(200, 500))

	// File explorer
	fileExplorerButton := widget.NewButton("Open explorer", func() {
		NewPathPicker(config.GetWindow(), func(path string) {
			// change user home dir to tilde (~)
			homeDir, _ := os.UserHomeDir()
			path = strings.Replace(path, homeDir, "~", 1)

			listedPaths = append(listedPaths, path)

			pathInput.SetText("")
			pathInput.Refresh()
			pathsList.Refresh()
		})
	})

	// Window content
	layout := container.NewBorder(
		container.NewVBox(
			sizeSelectorContainer,
			widget.NewSeparator(),
			pathsSelectorLabel,
			pathInput,
			fileExplorerButton,
		),
		nil, nil, nil,
		pathsList,
	)

	// Create the new dialog window (the main container)
	configDialog := dialog.NewCustomConfirm("🔧  Application Settings", "Confirm", "Cancel", layout,
		// When the user confirms the changes
		func(confirm bool) {
			if confirm {
				// update fyne config API
				config.SetGridSize(sizes[selectedGridSize])
				config.SetPaths(listedPaths)

				// refresh the main window
				callback()
			}
		}, config.GetWindow())

	configDialog.Resize(
		fyne.NewSize(
			config.GetWindow().Canvas().Size().Width*0.85,
			config.GetWindow().Canvas().Size().Height*0.85,
		),
	)

	configDialog.Show()
}
