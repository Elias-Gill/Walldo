package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Iconos predefinidos mapeados a constantes nativas de Fyne para evitar strings crudos.
const (
	IconSearch   = "search"
	IconRefresh  = "refresh"
	IconSettings = "settings"
	IconFolder   = "folder"
)

// NewIconButton crea un botón estilizado usando los iconos nativos del sistema/tema actual.
func NewIconButton(text string, iconName string, tapped func()) *widget.Button {
	var icon fyne.Resource

	switch iconName {
	case IconSearch:
		icon = theme.SearchIcon()
	case IconRefresh:
		icon = theme.ViewRefreshIcon()
	case IconSettings:
		icon = theme.SettingsIcon()
	case IconFolder:
		icon = theme.FolderOpenIcon()
	}

	if icon != nil {
		return widget.NewButtonWithIcon(text, icon, tapped)
	}
	return widget.NewButton(text, tapped)
}
