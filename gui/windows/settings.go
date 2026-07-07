package windows

import (
	"os"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/gui/components"
)

// Structural mappings for GridSize enums to UI strings
var sizeToString = map[config.GridSize]string{
	config.SMALL:  "Small",
	config.NORMAL: "Default",
	config.LARGE:  "Large",
}

var stringToSize = map[string]config.GridSize{
	"Small":   config.SMALL,
	"Default": config.NORMAL,
	"Large":   config.LARGE,
}

type SettingsWindow struct {
	parentWin fyne.Window
	onSave    func()
}

func NewSettingsWindow(parent fyne.Window, onSave func()) *SettingsWindow {
	return &SettingsWindow{parentWin: parent, onSave: onSave}
}

func (s *SettingsWindow) Show() {
	// 1. Grid Size Selector
	selectedGridSize := sizeToString[config.Config.GridSize]
	sizeSelector := widget.NewRadioGroup([]string{"Small", "Default", "Large"}, func(sel string) {
		selectedGridSize = sel
	})
	sizeSelector.SetSelected(selectedGridSize)

	// 2. Wallpaper Search Paths List
	listedPaths := append([]string{}, config.Config.WallpaperFolders...)
	pathsList := widget.NewList(
		func() int { return len(listedPaths) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) { o.(*widget.Label).SetText(listedPaths[i]) },
	)

	// Remove item on selection
	pathsList.OnSelected = func(id int) {
		listedPaths = append(listedPaths[:id], listedPaths[id+1:]...)
		pathsList.Refresh()
	}

	// 3. Path Manual Input
	pathInput := widget.NewEntry()
	if runtime.GOOS == "windows" {
		pathInput.SetPlaceHolder(`C:/Users/user/wallpapers`)
	} else {
		pathInput.SetPlaceHolder(`~/wallpapers`)
	}

	pathInput.OnSubmitted = func(t string) {
		if len(strings.TrimSpace(t)) > 0 {
			listedPaths = append(listedPaths, t)
			pathInput.SetText("")
			pathsList.Refresh()
		}
	}

	// 4. Native Directories Browser
	fileExplorer := components.NewIconButton("Open", components.IconFolder, func() {
		picker := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				homeDir, _ := os.UserHomeDir()
				path := strings.Replace(uri.Path(), homeDir, "~", 1)
				listedPaths = append(listedPaths, path)
				pathsList.Refresh()
			}
		}, s.parentWin)
		picker.Show()
	})

	// Assemble view hierarchy
	content := container.NewBorder(
		container.NewVBox(
			widget.NewLabel("Image grid cell size"),
			sizeSelector,
			widget.NewSeparator(),
			widget.NewLabel("Wallpapers paths (Click an item to remove it)"),
			container.New(&components.PriorityLayout{}, pathInput, fileExplorer),
		),
		nil, nil, nil,
		pathsList,
	)

	// Revert back to a Modal Dialog attached to the parent window context
	d := dialog.NewCustomConfirm("🔧 Application Settings", "Save Changes", "Cancel", content, func(confirm bool) {
		if confirm {
			// Persist states back to global configuration
			config.Config.GridSize = stringToSize[selectedGridSize]
			config.Config.WallpaperFolders = listedPaths

			config.PersistConfig()
			s.onSave()
		}
	}, s.parentWin)

	// Responsive scaling based on the trigger window dimension bounds
	d.Resize(fyne.NewSize(
		s.parentWin.Canvas().Size().Width*0.85,
		s.parentWin.Canvas().Size().Height*0.85,
	))
	d.Show()
}
