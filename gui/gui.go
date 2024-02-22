package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/elias-gill/walldo-in-go/globals"
	global "github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/gui/components"
)

func SetupGui() {
	// restore previous window size
	globals.RestoreWindowSize()

	// title
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	// scrollable image grid
	grid := components.NewImageGrid()

	// bottom nav
	bottomNav := components.NewBottomNav(grid.RefreshImgGrid)

	global.Window.SetContent(
		container.New(
			layout.NewBorderLayout(title, bottomNav, nil, nil),
			title,
			grid.GetGridContent(),
			bottomNav,
		),
	)

	// load images and thumbnails while initializing the GUI
	go grid.RefreshImgGrid()

	// save the window size on close
	global.MyApp.Lifecycle().SetOnStopped(func() {
		global.MyApp.Preferences().SetFloat("WindowHeight", float64(global.Window.Canvas().Size().Height))
		global.MyApp.Preferences().SetFloat("WindowWidth", float64(global.Window.Canvas().Size().Width))
	})
}
