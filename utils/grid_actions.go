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

// Fill the grid with images and refresh the container on each step
func SetNewContent(contenedor *fyne.Container) {
	listImagesRecursivelly() // search original images
	getResizedImages()

	contenedor.RemoveAll()
	for i := range globals.Original_images {
		cont := fillContainer(contenedor, i)

		// define grid format
		// TODO agregar un formato de lista
		switch globals.GridTitles {
		case "Borderless":
			contenedor.Add(cont)

		default:
			card := widget.NewCard("", isolateImageName(globals.Original_images[i]), cont)
			contenedor.Add(card)
		}
		contenedor.Refresh()
	}
}

// Creates the new component to push into the grid.
// Every component has a container, a button with the position of an image
// from the images list and a thumbnail
func fillContainer(fyneContainer *fyne.Container, i int) *fyne.Container {
	button := widget.NewButton(strconv.Itoa(i), nil)
	button.OnTapped = func() {
		value, _ := strconv.Atoi(button.Text)
		// the button has the index of the original image
		SetWallpaper(globals.Original_images[value])
	}
	// TODO  maybe this need a refactor
	// resize the image and get the thumbnail
	resizeImage(i)
	aux := canvas.NewImageFromFile(globals.Resized_images[i])
	aux.ScaleMode = canvas.ImageScaleFastest
	aux.FillMode = canvas.ImageFillContain

	// A bit of "magia" (With the max layout we can overlap the button and the thumbnail)
	cont := container.NewMax(aux, button)
	return cont
}

// Return a new size depending on the users config
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

// Returns a new Grid (this grid is an empty frame)
func NewContentGrid() (*fyne.Container, *fyne.Container) {
	content_grid := container.New(layout.NewGridWrapLayout(SetGridSize()))
	grid := container.New(layout.NewPaddedLayout(), container.NewScroll(content_grid)) // make the grid actually scrollable
	return grid, content_grid
}
