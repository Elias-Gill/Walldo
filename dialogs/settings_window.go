package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
)

// ventana de configuracion de la app
func ConfigWindow(win *fyne.Window, app fyne.App, refresh *widget.Button) {
	var sel_grid_style string
	var sel_grid_size string
	var sel_layout_style string

	// selector de tamano de la grilla
	grid_size_selector := widget.NewRadioGroup([]string{"large", "default", "small"},
		func(sel string) {
			sel_grid_size = sel
		})
	grid_size_selector.SetSelected(globals.GridSize)

	// selector de estilo de la grilla
	grid_style_selector := widget.NewRadioGroup(
		[]string{"Borderless", "Captions"},
		func(sel string) {
			sel_grid_style = sel
		})
	grid_style_selector.SetSelected(globals.GridTitles)

	// selector de estilo de layout
	layout_selector := widget.NewRadioGroup(
		[]string{"Grid", "Rows"},
		func(sel string) {
			sel_layout_style = sel
		})
	layout_selector.SetSelected(globals.LayoutStyle)

	// contenido del dialogo
	cont := []*widget.FormItem{
		widget.NewFormItem("Layout", layout_selector),
		widget.NewFormItem("Images size", grid_size_selector),
		widget.NewFormItem("Grid style", grid_style_selector),
	}

	// La forma menos elegante de implementar un boton de refresh, pero funciona XD
	dia := dialog.NewForm("Settings", "Confirm", "Cancel", cont,
		func(status bool) {
			if status {
				// actualizar la config
				globals.MyApp.Preferences().SetString("GridTitles", sel_grid_style)
				globals.MyApp.Preferences().SetString("GridSize", sel_grid_size)
				globals.MyApp.Preferences().SetString("Layout", sel_layout_style)

				// actualizar variables globales
				globals.GridTitles = sel_grid_style
				globals.GridSize = sel_grid_size
				globals.LayoutStyle = sel_layout_style
				// refrescar el programa
				refresh.OnTapped()
			}
		}, *win)

	dia.Show()
}
