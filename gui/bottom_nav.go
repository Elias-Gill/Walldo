package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

func NewBottomNav(refresh func()) *fyne.Container {
	// reload button (on the bottom right)
	refreshButton := NewButtonWithIcon("", refresh, ICON_REFRESH)

	// button to open the config menu
	configsButton := NewButtonWithIcon("Preferences", func() {
		newConfigWindow(refresh)
	}, ICON_SETTINGS)

	// fuzzy finder button
	fuzzyButton := NewButtonWithIcon("", func() {
		NewFuzzyDialog()
	}, ICON_SEARCH)

	// Convert []wallpaper.FillStyle to []string for the Fyne widget
	availableModes := wallpaper.AvailableModes()
	selectOptions := make([]string, len(availableModes))
	for i, mode := range availableModes {
		selectOptions[i] = ModeToStr(mode)
	}

	// scale mode selector
	strategySelector := widget.NewSelect(
		selectOptions,
		func(sel string) {
			config.SetWallpFillMode(StrToMode(sel))
		},
	)
	strategySelector.SetSelected(ModeToStr(config.GetWallpFillMode()))

	// assemble app layout
	bottomNav := container.New(layout.NewHBoxLayout(),
		strategySelector,
		fuzzyButton,
		layout.NewSpacer(), // Single spacer is enough to push buttons to the right
		refreshButton,
		configsButton,
	)

	return bottomNav
}

const (
	STRING_SCALE    = "Scale"
	STRING_TILE     = "Tile"
	STRING_CENTER   = "Center"
	STRING_ORIGINAL = "Original"
	STRING_ZOOM     = "Zoom fill"
)

var stringToMode = map[string]wallpaper.FillStyle{
	STRING_SCALE:    wallpaper.FILL_SCALE,
	STRING_TILE:     wallpaper.FILL_TILE,
	STRING_CENTER:   wallpaper.FILL_CENTER,
	STRING_ORIGINAL: wallpaper.FILL_ORIGINAL,
	STRING_ZOOM:     wallpaper.FILL_ZOOM,
}

var modeToString = map[wallpaper.FillStyle]string{
	wallpaper.FILL_SCALE:    STRING_SCALE,
	wallpaper.FILL_TILE:     STRING_TILE,
	wallpaper.FILL_CENTER:   STRING_CENTER,
	wallpaper.FILL_ORIGINAL: STRING_ORIGINAL,
	wallpaper.FILL_ZOOM:     STRING_ZOOM,
}

func StrToMode(s string) wallpaper.FillStyle {
	return stringToMode[s]
}

func ModeToStr(s wallpaper.FillStyle) string {
	return modeToString[s]
}
