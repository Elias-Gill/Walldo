package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	global "github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/gui/dialogs"
	"github.com/elias-gill/walldo-in-go/utils"
)

func SetupGui() {
	// instance a new fyne window and create a new layout
    c := wallpapersGrid{
        content: fyne.NewContainer()}
	c.defineLayout()

	// generate a new scrollable container for the body of the app
	mainFrame := container.New(
		layout.NewPaddedLayout(),
		container.NewScroll(c.content))

	// title style
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	// reload button (on the bottom right)
	refreshButton := newButton("", func() {
		// refresh the global variables, read the new config and reload thumbnails
		c.defineLayout()
		c.fillWithCards()
	}, "viewRefresh")

	// search bar with fuzzy finder
	fuzzyButton := newButton("", func() {
		dialogs.NewFuzzyDialog(global.Window)
	}, "search")

	// button with unsplash random image
	unsplashButton := newButton("", func() {
		utils.SetRandomImage()
	}, "mediaPhoto")

	// button that opens the config menu
	configsButton := newButton("Preferences", func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, func() {
			// refresh the global variables, read the new config and reload thumbnails
			c.defineLayout()
			c.fillWithCards()
		})
	}, "settings")

	// image scale mode selector
	strategySelector := widget.NewSelect([]string{"Zoom Fill", "Scale", "Center", "Original", "Tile"}, func(sel string) {
		global.FillStrategy = sel
		global.MyApp.Preferences().SetString("FillStrategy", sel)
	})
	// default selection
	strategySelector.SetSelected(global.FillStrategy)

	// setting the app content
	hbox := container.New(layout.NewHBoxLayout(),
		strategySelector,
		fuzzyButton,
		layout.NewSpacer(),
		unsplashButton,
		layout.NewSpacer(),
		refreshButton,
		configsButton,
	)
	content := container.New(layout.NewBorderLayout(title, hbox, nil, nil), title, mainFrame, hbox)
	global.Window.SetContent(content)

	// load images and thumbnails just after initializing the GUI
	global.MyApp.Lifecycle().SetOnStarted(func() {
		c.fillWithCards()
	})

	// save the window size on close
	global.Window.SetOnClosed(func() {
		println(global.Window.Canvas().Size().Height)
		println(global.Window.Canvas().Size().Width)
		global.MyApp.Preferences().SetFloat("WindowHeight", float64(global.Window.Canvas().Size().Height))
		global.MyApp.Preferences().SetFloat("WindowWidth", float64(global.Window.Canvas().Size().Width))
	})
}
