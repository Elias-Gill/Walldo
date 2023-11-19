package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
)

// Template for creating a new button with a custom icon.
func NewButton(text string, f func(), icon string) *widget.Button {
    if len(icon) > 0 {
        ico := fyne.ThemeIconName(icon)
        return widget.NewButtonWithIcon(text, globals.MyApp.Settings().Theme().Icon(ico), f)
    }
    return widget.NewButton(text, f)
}

const(
    ICON_SEARCH =  "search"
    ICON_REFRESH =  "viewRefresh"
    ICON_SETTINGS =  "settings"
)
