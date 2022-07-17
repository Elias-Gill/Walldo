package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/fuzzy_engine/matching"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
)

var data [5]string
var fuzzy_list *widget.List

// refresh the content search list with every keystroke
func entryChanged(entry string) {
	if len(entry) >= 1 {
		aux := matching.FindAll(entry, globals.Original_images)
		for i := 0; i < len(data); i++ {
            if i < len(aux){
                data[i] = globals.Original_images[aux[i].Idx]
            } else {
                data[i] = ""
            }
		}
	}
	fuzzy_list.Refresh()
}

// Create a new Fuzzy finder dialog and displays it
func NewFuzzyDialog(w fyne.Window) {
	// variables
	var fuzzy_searcher = widget.NewEntry()
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

	c := container.New(layout.NewBorderLayout(fuzzy_searcher, nil, nil, nil), fuzzy_searcher, fuzzy_list)
	x := dialog.NewCustom("Fuzzy searh", "Cancel", c, globals.Window)
	x.Resize(fyne.NewSize(500, 300))

	fuzzy_searcher.OnChanged = entryChanged
	fuzzy_searcher.SetPlaceHolder("Search Image")
	fuzzy_list.OnSelected = listSelected

    data = [5]string{}
	x.Show()
}

// change the wallpaper with the given selection
func listSelected(id int) {
	utils.SetWallpaper(data[id])
}
