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
	refresh_button := newButton("", func() {
		// actualizar configuracion y recargar imagenes
		grid_content.Layout = layout.NewGridWrapLayout(utils.SetGridSize())
		utils.SetNewContent(grid_content)

		grid.Refresh()
	}, "viewRefresh")

	// buscador de imagenes con algoritmo difuso
	fuzzy_button := newButton("", func() {
		dialogs.NewFuzzyDialog(global.Window)
	}, "search")

	// abrir el menu de configuraciones
	configs_button := newButton("Preferences", func() {
		dialogs.ConfigWindow(&global.Window, global.MyApp, refresh_button)
	}, "settings")

	// selector de modo de escalado de imagen
	strategy_selector := widget.NewSelect([]string{"Zoom Fill", "Scale", "Center", "Original", "Tile"}, func(sel string) {
		global.FillStrategy = sel
		global.MyApp.Preferences().SetString("FillStrategy", sel)
	})
	strategy_selector.SetSelected(global.FillStrategy)

	// contenido de la aplicacion
	hbox := container.New(layout.NewHBoxLayout(), strategy_selector, fuzzy_button, layout.NewSpacer(), refresh_button, configs_button)
	content := container.New(layout.NewBorderLayout(titulo, hbox, nil, nil), titulo, grid, hbox)
	global.Window.SetContent(content)

	// rellenar las imagenes solo despues de iniciar
	// corre en una go routine de manera concurrente (velocidad absurda)
	global.MyApp.Lifecycle().SetOnStarted(func() {
		go utils.SetNewContent(grid_content)
	})

	global.Window.ShowAndRun()
}

// retornar un nuevo boton con el icono especificado y la funcion especifica
func newButton(name string, f func(), icon ...string) *widget.Button {
	if len(icon) > 0 {
		ico := fyne.ThemeIconName(icon[0])
		return widget.NewButtonWithIcon(name, global.MyApp.Settings().Theme().Icon(ico), f)
	}
	return widget.NewButton(name, f)
}
