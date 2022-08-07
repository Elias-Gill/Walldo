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

	// instance a new fyne window and create a new grid layout
	global.Window.Resize(fyne.NewSize(1020, 600))
	grid, grid_content := utils.NewContentGrid()

	// main title
	titulo := canvas.NewText("Select your wallpaper", color.White)
	titulo.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 18

	// reload button (on the bottom right)
	refresh_button := newButton("", func() {
		// actualizar configuracion y recargar imagenes
		// refresh the global variables, read the new config and reload thumbnails
		grid_content.Layout = layout.NewGridWrapLayout(utils.SetGridSize())
		utils.SetNewContent(grid_content)

		grid.Refresh()
	}, "viewRefresh")

	// search bar with fuzzy finder
	fuzzy_button := newButton("", func() {
		dialogs.NewFuzzyDialog(global.Window)
	}, "search")

	// button that opens the config menu
	configs_button := newButton("Preferences", func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, refresh_button)
	}, "settings")

	// image scale mode selector
	strategy_selector := widget.NewSelect([]string{"Zoom Fill", "Scale", "Center", "Original", "Tile"}, func(sel string) {
		global.FillStrategy = sel
		global.MyApp.Preferences().SetString("FillStrategy", sel)
	})
	strategy_selector.SetSelected(global.FillStrategy)

	// setting the app content
	hbox := container.New(layout.NewHBoxLayout(), strategy_selector, fuzzy_button, layout.NewSpacer(), refresh_button, configs_button)
	content := container.New(layout.NewBorderLayout(titulo, hbox, nil, nil), titulo, grid, hbox)
	global.Window.SetContent(content)

	// load images and thumbnails concurrently just after initializing the GUI
	// to improve user experience.
	global.MyApp.Lifecycle().SetOnStarted(func() {
		go utils.SetNewContent(grid_content)
	})

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
