package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
)

const (
	ICON_SEARCH   = "search"
	ICON_REFRESH  = "viewRefresh"
	ICON_SETTINGS = "settings"
)

// Template for creating a new button with a custom icon.
func NewButtonWithIcon(app *globals.App, text string, icon string, action func()) *widget.Button {
	if len(icon) > 0 {
		ico := fyne.ThemeIconName(icon)
		return widget.NewButtonWithIcon(text, app.App.Settings().Theme().Icon(ico), action)
	}

	return widget.NewButton(text, action)
}
