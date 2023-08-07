package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/fuzzyEngine/matching"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

var (
	data       = []string{}
	fuzzy_list *widget.List
)

// refresh the content search list with every keystroke
func entryChanged(entry string) {
	data = []string{}
	imagesList := utils.GetImagesList()
	if len(entry) >= 1 {
		// search for the matching results
		matches := matching.FindAll(entry, imagesList)
		// display the results
		for i := 0; i < len(matches); i++ {
			data = append(data, imagesList[matches[i].Idx])
		}
	}
	fuzzy_list.Refresh()
}

// Create a new Fuzzy finder dialog and displays it
func NewFuzzyDialog(w fyne.Window) {
	// variables
	fuzzy_searcher := widget.NewEntry()
	// list of results
	fuzzy_list = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	cont := container.New(layout.NewBorderLayout(fuzzy_searcher, nil, nil, nil), fuzzy_searcher, fuzzy_list)
	dial := dialog.NewCustom("Fuzzy search", "Cancel", cont, globals.Window)
	dial.Resize(fyne.NewSize(500, 300))

	// opciones
	fuzzy_searcher.OnChanged = entryChanged
	fuzzy_searcher.SetPlaceHolder("Search Image")
	fuzzy_list.OnSelected = func(id int) {
		wallpaper.SetWallpaper(data[id])
	}

	dial.Show()
}
