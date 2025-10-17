package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
)

// Template for creating a new button with a custom icon.
func NewButtonWithIcon(text string, f func(), icon string) *widget.Button {
	if len(icon) > 0 {
		ico := fyne.ThemeIconName(icon)
		return widget.NewButtonWithIcon(text, config.GetFyneSettings().Theme().Icon(ico), f)
	}

	return widget.NewButton(text, f)
}

const (
	ICON_SEARCH   = "search"
	ICON_REFRESH  = "viewRefresh"
	ICON_SETTINGS = "settings"
	ICON_FOLDER   = "folder"
)
