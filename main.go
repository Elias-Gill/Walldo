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
	// settear todas las variables globales dependiendo del OS
	global.SetGlobalValues()

	// instanciar la nueva ventana
	global.Window.Resize(fyne.NewSize(950, 600))
	var content *fyne.Container
	// generar la grilla de imagenes
	grid, grid_content := utils.NewContentGrid()

	// titulo principal
	titulo := canvas.NewText("Select your wallpaper", color.White)
	titulo.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 18

	// reload button
	ico := fyne.ThemeIconName("viewRefresh")
	refresh_button := widget.NewButtonWithIcon("", global.MyApp.Settings().Theme().Icon(ico), func() {
		// actualizar configuracion y recargar imagenes
		grid_content.Layout = layout.NewGridWrapLayout(utils.SetGridSize())
		utils.SetNewContent(grid_content)

		grid.Refresh()
	})

	// buscador de imagenes con algoritmo difuso
	ico = fyne.ThemeIconName("search")
	fuzzy_button := widget.NewButtonWithIcon("", global.MyApp.Settings().Theme().Icon(ico), func() {
		dialogs.NewFuzzyDialog(global.Window)
	})

	// abrir el menu de configuraciones
	ico = fyne.ThemeIconName("settings")
	configs_button := widget.NewButtonWithIcon("Preferences", global.MyApp.Settings().Theme().Icon(ico), func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, refresh_button)
	})

	// selector de modo de escalado de imagen
	strategy_selector := widget.NewSelect([]string{"Zoom Fill", "Scale", "Center", "Original", "Tile"}, func(sel string) {
		global.FillStrategy = sel
		global.MyApp.Preferences().SetString("FillStrategy", sel)
	})
	strategy_selector.SetSelected(global.FillStrategy)

	hbox := container.New(layout.NewHBoxLayout(), strategy_selector, fuzzy_button, layout.NewSpacer(), refresh_button, configs_button)
	content = container.New(layout.NewBorderLayout(titulo, hbox, nil, nil), titulo, grid, hbox)
	global.Window.SetContent(content)

	// rellenar las imagenes solo despues de iniciar
	// corre en una go routine de manera concurrente
	global.MyApp.Lifecycle().SetOnStarted(func() {
		go utils.SetNewContent(grid_content)
	})

	global.Window.ShowAndRun()
}
