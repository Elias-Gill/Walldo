package dialogs

/* import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	global "github.com/elias-gill/walldo-in-go/globals"
) */

/* func DisplayError(err error) {
    app := global.MyApp
	cont := []*widget.FormItem{
		// selector de tamano de la grilla
		widget.NewFormItem("Tamano de \n imagenes", widget.NewSelect(
			[]string{"large", "default", "small"},
			func(sel string) {
				app.Preferences().SetString("gridSize", sel)
			})),

		// selector de tamano estilo de imagenes
		widget.NewFormItem("Estilo de grilla", widget.NewSelect(
			[]string{"borderless", "default"},
			func(sel string) {
				app.Preferences().SetString("layout", sel)
			})),
	}

	// La forma menos elegante de implementar un boton de refresh, pero funciona XD
	dia := dialog.NewForm("Settings", "Confirm", "Cancel", cont, func(bool) { refresh.OnTapped() }, *win)
	dia.Show()
} */
