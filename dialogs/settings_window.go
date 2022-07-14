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

    // selector de tamano de la grilla
    grid_size_selector := widget.NewRadioGroup( []string{"large", "default", "small"},
        func(sel string) {
            sel_grid_size = sel
    })
    grid_size_selector.SetSelected(globals.GridSize)

    // selector de estilo de la grilla
    grid_style_selector := widget.NewRadioGroup(
        []string{"borderless", "default"},
        func(sel string) {
            sel_grid_style = sel
    })
    grid_style_selector.SetSelected(globals.GridTitles)

    // contenido del dialogo
    cont := []*widget.FormItem{
        widget.NewFormItem("Tamano de \n imagenes", grid_size_selector),
        widget.NewFormItem("Estilo de grilla", grid_style_selector),
    }

	// La forma menos elegante de implementar un boton de refresh, pero funciona XD
	dia := dialog.NewForm("Settings", "Confirm", "Cancel", cont, 
        func(status bool) { 
            if(status){
                // actualizar la config
                globals.MyApp.Preferences().SetString("layout", sel_grid_style)
                globals.MyApp.Preferences().SetString("gridSize", sel_grid_size)

                // actualizar variables globales
                globals.GridTitles = sel_grid_style
                globals.GridSize = sel_grid_size
                // refrescar el programa
                refresh.OnTapped() 
            }
        }, *win)

	dia.Show()
}
