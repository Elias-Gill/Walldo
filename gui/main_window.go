package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	global "github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/gui/components"
	"github.com/elias-gill/walldo-in-go/gui/components/dialogs"
)

func SetupGui() {
	// title style
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	// generate a new scrollable that contains the images grid
	grid := components.NewImageGrid()
	mainFrame := container.New(
		layout.NewPaddedLayout(),
		container.NewScroll(grid.GetGridContent()))

	// reload button (on the bottom right)
	refreshButton := components.NewButtonWithIcon("", grid.RefreshImgGrid, components.ICON_REFRESH)

	// button to open the config menu
	configsButton := components.NewButtonWithIcon("Preferences", func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, grid.RefreshImgGrid)
	}, components.ICON_SETTINGS)

	// fuzzy finder button
	fuzzyButton := components.NewButtonWithIcon("", func() {
		dialogs.NewFuzzyDialog(global.Window)
	}, components.ICON_SEARCH)

	// scale mode selector
	strategySelector := widget.NewSelect(
		[]string{
			global.FILL_ZOOM, global.FILL_TILE, global.FILL_SCALE, global.FILL_CENTER, global.FILL_ORIGINAL,
		},
		func(sel string) {
			global.FillStrategy = sel
		},
	)
	strategySelector.SetSelected(global.FillStrategy)

	// assemble app layout
	body := container.New(layout.NewHBoxLayout(),
		strategySelector,
		fuzzyButton,
		layout.NewSpacer(),
		// FUTURE: imageName,
		layout.NewSpacer(),
		refreshButton,
		configsButton,
	)
	content := container.New(layout.NewBorderLayout(title, body, nil, nil), title, mainFrame, body)
	global.Window.SetContent(content)

	// load images and thumbnails just after initializing the GUI
	global.MyApp.Lifecycle().SetOnStarted(func() {
		grid.RefreshImgGrid()
	})

	// save the window size on close
	global.MyApp.Lifecycle().SetOnStopped(func() {
		global.MyApp.Preferences().SetFloat("WindowHeight", float64(global.Window.Canvas().Size().Height))
		global.MyApp.Preferences().SetFloat("WindowWidth", float64(global.Window.Canvas().Size().Width))
	})
}
