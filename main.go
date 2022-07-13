package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	dialogs "github.com/elias-gill/walldo-in-go/dialogs"
	"github.com/elias-gill/walldo-in-go/utils"
)

func main() {
	// instanciar la nueva ventana
	myApp := app.NewWithID("walldo")
	w := myApp.NewWindow("Walldo in go")
	w.Resize(fyne.NewSize(800, 500))
	var content *fyne.Container

	// configuracion de usuario (parte estetica)
	gridSize := myApp.Preferences().StringWithFallback("gridSize", "default")
	layoutStyle := myApp.Preferences().StringWithFallback("layout", "default")

	// generar la grilla de imagenes
	grid, grid_content := utils.NewContentGrid(gridSize)

	// titulo de la app
	titulo := canvas.NewText("Select your wallpaper", color.White)
	titulo.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 18

	// botones principales
	refresh_button := widget.NewButton("Reload", func() {
		// actualizar configuracion y recargar imagenes
		gridSize = myApp.Preferences().StringWithFallback("gridSize", "default")
		layoutStyle = myApp.Preferences().StringWithFallback("layout", "default")

		grid_content.Layout = layout.NewGridWrapLayout(utils.SetGridSize(gridSize))
		utils.SetNewContent(grid_content, layoutStyle)

		grid.Refresh()
	})

	// abrir el menu de configuraciones
	configs_button := widget.NewButton("Preferences", func() {
		dialogs.ConfigWindow(&w, myApp, refresh_button)
	})

	hbox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), refresh_button, configs_button)
	content = container.New(layout.NewBorderLayout(titulo, hbox, nil, nil), titulo, grid, hbox)
	w.SetContent(content)

	// rellenar las imagenes solo despues de iniciar
	// corre en una go routine de manera concurrente
	myApp.Lifecycle().SetOnStarted(func() {
		go utils.SetNewContent(grid_content, layoutStyle)
	})

	w.ShowAndRun()
}
