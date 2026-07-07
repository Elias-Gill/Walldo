package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/gui/themes"
	"github.com/elias-gill/walldo-in-go/gui/windows"
)

func StartGui() {
	myApp := app.NewWithID("Walldo")
	myApp.Settings().SetTheme(themes.NewDarkTheme())

	// Instanciamos la ventana principal orientada a objetos
	home := windows.NewHomeWindow(myApp)
	win := home.Window()

	// Guardar estado al cerrar de manera limpia
	myApp.Lifecycle().SetOnStopped(func() {
		myApp.Preferences().SetFloat("WindowHeight", float64(win.Canvas().Size().Height))
		myApp.Preferences().SetFloat("WindowWidth", float64(win.Canvas().Size().Width))
		config.PersistConfig()
	})

	// Recuperar tamaño previo
	win.Resize(fyne.NewSize(
		float32(myApp.Preferences().FloatWithFallback("WindowWidth", 800)),
		float32(myApp.Preferences().FloatWithFallback("WindowHeight", 800)),
	))

	win.ShowAndRun()
}
