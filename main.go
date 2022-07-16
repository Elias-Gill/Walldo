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
    global.Window.Resize(fyne.NewSize(800, 500))
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

	// botones principales
	refresh_button := widget.NewButton("Reload", func() {
		// actualizar configuracion y recargar imagenes
		grid_content.Layout = layout.NewGridWrapLayout(utils.SetGridSize())
		utils.SetNewContent(grid_content)

		grid.Refresh()
	})

	// abrir el menu de configuraciones
	configs_button := widget.NewButton("Preferences", func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, refresh_button)
	})

    // selector de modo de escalado de imagen
    strategy_selector := widget.NewSelect([]string{"Zoom Fill", "Scale", "Center", "Original", "Tile"}, func(sel string) {
        global.FillStrategy = sel
        global.MyApp.Preferences().SetString("FillStrategy", sel)
	})
    strategy_selector.SetSelected(global.FillStrategy)

	hbox := container.New(layout.NewHBoxLayout(), strategy_selector,layout.NewSpacer(), refresh_button, configs_button)
	content = container.New(layout.NewBorderLayout(titulo, hbox, nil, nil), titulo, grid, hbox)
	global.Window.SetContent(content)

	// rellenar las imagenes solo despues de iniciar
	// corre en una go routine de manera concurrente
	global.MyApp.Lifecycle().SetOnStarted(func() {
		go utils.SetNewContent(grid_content)
	})

	global.Window.ShowAndRun()
}
