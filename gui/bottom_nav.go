package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/wallpaper"
	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

func NewBottomNav(refresh func()) *fyne.Container {
	// reload button (on the bottom right)
	refreshButton := NewButtonWithIcon("", refresh, ICON_REFRESH)

	// button to open the config menu
	configsButton := NewButtonWithIcon("Preferences", func() {
		newConfigWindow(refresh)
	}, ICON_SETTINGS)

	// fuzzy finder button
	// fuzzyButton := NewButtonWithIcon("", func() {
	// 	dialogs.NewFuzzyDialog()
	// }, ICON_SEARCH)

	// scale mode selector
	strategySelector := widget.NewSelect(
		wallpaper.ListModes(),
		func(sel string) {
			config.SetWallpFillMode(modes.StrToMode(sel))
		})
	strategySelector.SetSelected(modes.ModeToStr(config.GetWallpFillMode()))

	// assemble app layout
	bottomNav := container.New(layout.NewHBoxLayout(),
		strategySelector,
		// fuzzyButton,
		layout.NewSpacer(),
		layout.NewSpacer(),
		refreshButton,
		configsButton)

	return bottomNav
}
