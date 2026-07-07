package windows

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/fuzzyEngine/matching"
	"github.com/elias-gill/walldo-in-go/utils"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

func NewFuzzyDialog(parent fyne.Window) {
	var data []string

	resultsWidget := widget.NewList(
		func() int { return len(data) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(id widget.ListItemID, canvas fyne.CanvasObject) {
			canvas.(*widget.Label).SetText(data[id])
		},
	)

	resultsWidget.OnSelected = func(id int) {
		wallpaper.SetWallpaper(strings.Clone(data[id]), config.Config.WallpfillMode)
	}

	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Search Image")
	searchInput.OnChanged = func(entry string) {
		var imagesList []string
		for _, v := range utils.ListImages() {
			imagesList = append(imagesList, v.Path)
		}

		data = []string{}

		if len(entry) >= 1 {
			matches := matching.FindAll(entry, imagesList)
			for i := range matches {
				data = append(data, imagesList[matches[i].Idx])
			}
		}
		resultsWidget.Refresh()
	}

	content := container.NewBorder(searchInput, nil, nil, nil, searchInput, resultsWidget)

	dial := dialog.NewCustom("Fuzzy search", "Cancel", content, parent)
	dial.Resize(fyne.NewSize(500, 300))
	dial.Show()
}
