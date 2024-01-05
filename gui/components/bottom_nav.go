package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	global "github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/gui/components/dialogs"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

func NewBottomNav(refresh func()) *fyne.Container {
	// reload button (on the bottom right)
	refreshButton := NewButtonWithIcon("", refresh, ICON_REFRESH)

	// button to open the config menu
	configsButton := NewButtonWithIcon("Preferences", func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, refresh)
	}, ICON_SETTINGS)

	// fuzzy finder button
	fuzzyButton := NewButtonWithIcon("", func() {
		dialogs.NewFuzzyDialog()
	}, ICON_SEARCH)

	// scale mode selector
	strategySelector := widget.NewSelect(
		wallpaper.ListAvailableModes(),
		func(sel string) {
			global.FillStrategy = sel
		})
	strategySelector.SetSelected(global.FillStrategy)

	// assemble app layout
	bottomNav := container.New(layout.NewHBoxLayout(),
		strategySelector,
		fuzzyButton,
		layout.NewSpacer(),
		// FUTURE: imageName,
		layout.NewSpacer(),
		refreshButton,
		configsButton)

	return bottomNav
}
