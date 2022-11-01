package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	dialogs "github.com/elias-gill/walldo-in-go/dialogs"
	global "github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
)

func main() {
	// set all global variables for the instance
	global.SetGlobalValues()

	// instance a new fyne window and create a new layout
	global.Window.Resize(fyne.NewSize(1020, 600))
	mainContent := fyne.NewContainer()
	mainContent.Layout = utils.DefineLayout()
	mainContainer := container.New(layout.NewPaddedLayout(), container.NewScroll(mainContent)) // make the container scrollable

	// main title
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	// reload button (on the bottom right)
	refreshButton := newButton("", func() {
		// actualizar configuracion y recargar imagenes
		// refresh the global variables, read the new config and reload thumbnails
		mainContent.Layout = utils.DefineLayout()
		utils.CompleteCards(mainContent)
		mainContainer.Refresh()
	}, "viewRefresh")

	// search bar with fuzzy finder
	fuzzyButton := newButton("", func() {
		dialogs.NewFuzzyDialog(global.Window)
	}, "search")

	// button that opens the config menu
	configsButton := newButton("Preferences", func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, refreshButton)
	}, "settings")

	// image scale mode selector
	strategySelector := widget.NewSelect([]string{"Zoom Fill", "Scale", "Center", "Original", "Tile"}, func(sel string) {
		global.FillStrategy = sel
		global.MyApp.Preferences().SetString("FillStrategy", sel)
	})
	strategySelector.SetSelected(global.FillStrategy)

	// setting the app content
	hbox := container.New(layout.NewHBoxLayout(), strategySelector, fuzzyButton, layout.NewSpacer(), refreshButton, configsButton)
	content := container.New(layout.NewBorderLayout(title, hbox, nil, nil), title, mainContainer, hbox)
	global.Window.SetContent(content)

	// load images and thumbnails concurrently just after initializing the GUI
	// to improve user experience.
	global.MyApp.Lifecycle().SetOnStarted(func() {
		utils.CompleteCards(mainContent)
	})

	// run app
	global.Window.ShowAndRun()
}

// template for creating a new button with the specified function and icon name
func newButton(name string, f func(), icon ...string) *widget.Button {
	if len(icon) > 0 {
		ico := fyne.ThemeIconName(icon[0])
		return widget.NewButtonWithIcon(name, global.MyApp.Settings().Theme().Icon(ico), f)
	}
	return widget.NewButton(name, f)
}
