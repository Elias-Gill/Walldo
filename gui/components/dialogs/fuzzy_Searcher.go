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
	data          = []string{}
	resultsWidget *widget.List
)

// refresh the content search list with every keystroke.
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

	resultsWidget.Refresh()
}

// Create a new Fuzzy finder dialog and display it.
func NewFuzzyDialog(w fyne.Window) {
	// searcher configuration
	searcherWiget := widget.NewEntry()
	searcherWiget.SetPlaceHolder("Search Image")
	searcherWiget.OnChanged = entryChanged
	resultsWidget.OnSelected = func(id int) {
		wallpaper.SetFromFile(data[id])
	}

	// list of results widget
	resultsWidget = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	cont := container.New(layout.NewBorderLayout(searcherWiget, nil, nil, nil), searcherWiget, resultsWidget)
	dial := dialog.NewCustom("Fuzzy search", "Cancel", cont, globals.Window)
	dial.Resize(fyne.NewSize(500, 300))
	dial.Show()
}
