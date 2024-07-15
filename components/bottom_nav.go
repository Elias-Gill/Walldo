package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/components/dialogs"
	global "github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

func NewBottomNav(app *global.App, action func()) *fyne.Container {
	// reload button (on the bottom right)
	refreshButton := NewButtonWithIcon(app, "", ICON_REFRESH, action)

	// button to open the config menu
	configsButton := NewButtonWithIcon(app, "Preferences", ICON_SETTINGS, func() {
		dialogs.ConfigWindow(app, action)
	})

	// fuzzy finder button
	fuzzyButton := NewButtonWithIcon(app, "", ICON_SEARCH, func() {
		dialogs.NewFuzzyDialog(app)
	})

	// scale mode selector
	strategySelector := widget.NewSelect(
		wallpaper.ListAvailableModes(),
		func(sel string) {
			app.Config.FillStrategy = global.FillStyle(sel)
			app.WriteConfig()
		},
	)
	strategySelector.SetSelected(string(app.Config.FillStrategy))

	// assemble app layout
	bottomNav := container.New(layout.NewHBoxLayout(),
		strategySelector,
		fuzzyButton,
		layout.NewSpacer(),
		layout.NewSpacer(),
		refreshButton,
		configsButton)

	return bottomNav
}
