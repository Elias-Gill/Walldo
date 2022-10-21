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
)

var (
	data       = []string{}
	fuzzy_list *widget.List
)

// refresh the content search list with every keystroke
func entryChanged(entry string) {
	data = []string{}
	if len(entry) >= 1 {
		// search for the matching results
		aux := matching.FindAll(entry, globals.OriginalImages)
		// display the results
		for i := 0; i < len(aux); i++ {
			data = append(data, globals.OriginalImages[aux[i].Idx])
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
	fuzzy_list.OnSelected = listSelected

	dial.Show()
}

// change the wallpaper with the given selection
func listSelected(id int) {
	utils.SetWallpaper(data[id])
}
