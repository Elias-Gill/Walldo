package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	explorer "github.com/elias-gill/walldo-in-go/file_explorer"
	"github.com/elias-gill/walldo-in-go/utils"
)

func main() {
	// instanciar la nueva ventana
	myApp := app.NewWithID("walldo")
	w := myApp.NewWindow("Walldo in go")
	w.Resize(fyne.NewSize(800, 500))

	// generar la grilla de imagenes
	grid, grid_content := utils.NewContentGrid(&myApp)

	// titulo de la app
	titulo := canvas.NewText("Wallpapers with Go", color.White)
	titulo.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 18

	// botones
	refresh_button := widget.NewButton("Restore", func() {
        utils.SetGridContent(grid_content)
	})

	configs_button := widget.NewButton("Preferences", func() {
		// abrir el menu de configuraciones
		explorer.ConfigWindow(&w)
	})

	hbox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), refresh_button, configs_button)
	content := container.New(layout.NewBorderLayout(titulo, hbox, grid, nil), titulo, grid, hbox)
	w.SetContent(content)
	w.SetFixedSize(true)

    // rellenar las imagenes solo despues de iniciar
    // corre en una go routine para poder mostrar el menu grafico a la vez 
    // que se generan las tumbnails
    myApp.Lifecycle().SetOnStarted(func(){
        go utils.SetGridContent(grid_content) 
    })

	w.ShowAndRun()
}
