package gui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/fuzzyEngine/matching"
	"github.com/elias-gill/walldo-in-go/utils"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

// Create a new Fuzzy finder dialog and display it.
func NewFuzzyDialog() {
	data := []string{}

	// list of results
	resultsWidget := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, canvas fyne.CanvasObject) {
			canvas.(*widget.Label).SetText(data[id])
		})

	resultsWidget.OnSelected = func(id int) {
		wallpaper.SetWallpaper(strings.Clone(data[id]))
	}

	// search input
	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Search Image")
	searchInput.OnChanged =
		func(entry string) {
			var imagesList []string

			for _, v := range utils.ListImages() {
				imagesList = append(imagesList, v.Path)
			}

			data = []string{}

			if len(entry) >= 1 {
				// search for the matching results
				matches := matching.FindAll(entry, imagesList)
				// display the results
				for i := 0; i < len(matches); i++ {
					data = append(data, imagesList[matches[i].Idx])
				}
			}

			resultsWidget.Refresh()
		}

	cont := container.New(layout.NewBorderLayout(searchInput, nil, nil, nil), searchInput, resultsWidget)
	dial := dialog.NewCustom("Fuzzy search", "Cancel", cont, config.GetWindow())
	dial.Resize(fyne.NewSize(500, 300))
	dial.Show()
}
