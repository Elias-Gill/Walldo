package utils

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
)

var original_images []string
var resized_images []string

// rellenar la grilla de imagenes de manera asincrona y utilizando
// concurrencia
func SetNewContent(contenedor *fyne.Container) {
	listarImagenes()                   // buscar las imagenes
	getResizedImages(&original_images) // images reescaladas

	contenedor.RemoveAll()
	for i := range original_images {
		cont := rellenarContenedor(contenedor, i)

		// definir el formato de la aplicacion
		// TODO agregar un formato de lista
		switch globals.GridTitles {
		case "Borderless":
			contenedor.Add(cont)

		default:
			card := widget.NewCard("", aislarNombreImagen(original_images[i]), cont)
			contenedor.Add(card)
		}
		contenedor.Refresh()
	}
}

// Crea un elemento para la grilla de imagenes y lo anade a la grilla
// Cada imagen tiene asignado un boton, boton que contiene como texto la posicion de dicha imagen
// en el arreglo de imagenes originales
// Al acabar refresca el contenido
func rellenarContenedor(contenedor *fyne.Container, i int) *fyne.Container {
	button := widget.NewButton(strconv.Itoa(i), nil)
	button.OnTapped = func() {
		value, _ := strconv.Atoi(button.Text)
		// el boton contiene el index de la imagen original
		SetWallpaper(original_images[value])
	}
	// imagen rescalada
	resizeImage(i)
	aux := canvas.NewImageFromFile(resized_images[i])
	aux.ScaleMode = canvas.ImageScaleFastest
	aux.FillMode = canvas.ImageFillContain

	// algo de magia (el boton se le superpone a la imagen)
	cont := container.NewMax(aux, button)
	return cont
}

// retorna un tamano dependiendo de la configuracion del usuario
func SetGridSize() fyne.Size {
	switch globals.GridSize {
	case "small":
		return fyne.NewSize(110, 100)
	case "large":
		return fyne.NewSize(195, 175)
	default:
		return fyne.NewSize(150, 130)
	}
}

// Retorna la grilla de imagenes a ser mostradas.
func NewContentGrid() (*fyne.Container, *fyne.Container) {
	content_grid := container.New(layout.NewGridWrapLayout(SetGridSize()))
	grid := container.New(layout.NewPaddedLayout(), container.NewScroll(content_grid)) // make the grid actually scrollable
	return grid, content_grid
}
