package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/gui"
)

func main() {
	app := app.NewWithID("Walldo")
	window := app.NewWindow("Walldo")

	config.SetFyneSettings(app.Settings())
	config.SetWindow(window)

	// title
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	// scrollable image grid
	grid := gui.NewImageGrid()

	// bottom nav with buttons
	nav := gui.NewBottomNav(grid.RefreshImgGrid)

	window.SetContent(
		container.New(
			layout.NewBorderLayout(title, nav, nil, nil),
			title,
			grid.GetContent(),
			nav,
		),
	)

	// load images and thumbnails while initializing the GUI
	go grid.RefreshImgGrid()

	// save the window size on close
	app.Lifecycle().SetOnStopped(func() {
		app.Preferences().SetFloat("WindowHeight", float64(window.Canvas().Size().Height))
		app.Preferences().SetFloat("WindowWidth", float64(window.Canvas().Size().Width))
		config.WriteConfig()
	})

	// restore previous window size
	window.Resize(
		fyne.NewSize(
			float32(app.Preferences().FloatWithFallback("WindowWidth", 800)),
			float32(app.Preferences().FloatWithFallback("WindowHeight", 800)),
		),
	)

	window.ShowAndRun()
}
