package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/elias-gill/walldo-in-go/components"
	"github.com/elias-gill/walldo-in-go/globals"
)

func main() {
	app := globals.NewApp()

	// title
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	grid := components.NewImageGrid(app)
	bottomNav := components.NewBottomNav(app, grid.RefreshImgGrid)

	app.Window.SetContent(
		container.New(
			layout.NewBorderLayout(title, bottomNav, nil, nil),
			title,
			grid.GetGridContent(),
			bottomNav,
		),
	)

	// load images and thumbnails while initializing the GUI
	go grid.RefreshImgGrid()

	app.Window.ShowAndRun()
}
