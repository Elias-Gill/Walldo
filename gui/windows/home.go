package windows

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/gui/components"
)

type HomeWindow struct {
	win     fyne.Window
	gallery *components.WallpaperGallery
}

func NewHomeWindow(app fyne.App) *HomeWindow {
	h := &HomeWindow{
		win:     app.NewWindow("Walldo"),
		gallery: components.NewGallery(),
	}
	h.setupUI()
	return h
}

func (h *HomeWindow) setupUI() {
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	nav := h.buildBottomNav()

	h.win.SetContent(
		container.New(
			layout.NewBorderLayout(title, nav, nil, nil),
			title,
			h.gallery.View(),
			nav,
		),
	)

	// Async initial load for images
	go h.gallery.RefreshGallery()
}

func (h *HomeWindow) buildBottomNav() *fyne.Container {
	refreshBtn := components.NewIconButton("", components.IconRefresh, func() {
		go h.gallery.RefreshGallery()
	})

	configBtn := components.NewIconButton("Preferences", components.IconSettings, func() {
		settings := NewSettingsWindow(h.win, func() {
			h.gallery.RefreshGallery()
		})
		settings.Show()
	})

	strategySelector := widget.NewSelect([]string{"Scale", "Tile", "Center", "Zoom fill"}, func(sel string) {
		config.Config.WallpfillMode = components.StrToMode(sel)
	})
	strategySelector.SetSelected(components.ModeToStr(config.Config.WallpfillMode))

	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search...")
	searchBar.OnChanged = func(entry string) {
		h.gallery.FilterGallery(entry)
	}

	searchIcon := widget.NewIcon(theme.SearchIcon())
	searchGroup := container.NewBorder(nil, nil, searchIcon, nil, searchBar)

	// Wrap searchGroup to enforce a fixed optimal width inside the box layout
	searchWrapper := container.NewGridWrap(fyne.NewSize(250, searchBar.MinSize().Height), searchGroup)

	rightButtons := container.NewHBox(refreshBtn, configBtn)

	// Use Spacers to keep the fixed-width search bar perfectly centered
	return container.NewHBox(
		strategySelector,
		layout.NewSpacer(),
		searchWrapper,
		layout.NewSpacer(),
		rightButtons,
	)
}

func (h *HomeWindow) Window() fyne.Window {
	return h.win
}
