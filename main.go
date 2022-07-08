package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/utils"
	"image/color"
)

func main() {
	// instanciar la nueva ventana
	myApp := app.New()
	w := myApp.NewWindow("Walldo in go")
	w.Resize(fyne.NewSize(800, 500))

	// generar la grilla de imagenes
	grid := utils.NewContentGrid()

	// titulo de la app
	titulo := canvas.NewText("Wallpapers with Go", color.White)
	titulo.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 18

	// botones
	refresh_button := widget.NewButton("Restore", func() {
		// utils.Config_preferences()
		print("No implementado")
	})

	configs_button := widget.NewButton("Preferences", func() {
		// utils.Config_preferences()
		print("No implementado")
	})

	hbox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), refresh_button, configs_button)
	content := container.New(layout.NewBorderLayout(titulo, hbox, grid, nil), titulo, grid, hbox)
	w.SetContent(content)
	w.SetFixedSize(true)
	w.ShowAndRun()
}
