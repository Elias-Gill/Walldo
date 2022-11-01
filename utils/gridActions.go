// TODO: refactor the logic for thumbnails and loading
package utils

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
)

// Fill the grid with images and refresh the container on each step
func DefineLayout() fyne.Layout {
	// maybe in the future i can add more layouts
	return layout.NewGridWrapLayout(SetGridSize())
}

// fills the container with the correspondent content
func CompleteCards(c *fyne.Container) {
	c.RemoveAll()
	listImagesRecursivelly() // search original images
	getResizedImages()

	// save all images into a go channel to manage concurrently load/generate thumbnails
	channel := make(chan int, len(globals.OriginalImages))
	for i := range globals.OriginalImages {
		channel <- i
	}

	// create more "threads" to increase performance
	for i := 0; i < runtime.NumCPU()-2; i++ {
		go addNewCard(channel, c)
	}
	print("\n Usando ", runtime.NumCPU()-2, " Hilos")
}

// recibes the channel with the list of images and creates a new card of
func addNewCard(chanel chan int, c *fyne.Container) {
	for i := range chanel {
		content := generateFyneContent(i)

		switch globals.GridTitles {
		// grid without captions
		case "Borderless":
			c.Add(content)

		// normal grid with captions
		default:
			card := widget.NewCard("", isolateImageName(globals.OriginalImages[i]), content)
			c.Add(card)
		}
		c.Refresh()
	}
}

// Creates the new component to push into the grid.
// Every component has a container, a button with the position of an image
// from the images list and a thumbnail
func generateFyneContent(i int) *fyne.Container {
	button := widget.NewButton("", nil)
	button.OnTapped = func() {
		// the button has the index of the original image
		SetWallpaper(globals.OriginalImages[i])
	}

	// resize the image and get the thumbnail
	resizeImage(i)
	aux := canvas.NewImageFromFile(globals.ResizedImages[i])
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
