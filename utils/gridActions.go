package utils

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

// Generates and return a new layout acording to the user configurations
func DefineLayout() fyne.Layout {
	// default card size
	size := fyne.NewSize(150, 130)
	// other grid sizes
	switch globals.GridSize {
	case "small":
		size = fyne.NewSize(110, 100)
	case "large":
		size = fyne.NewSize(195, 175)
	}
	return layout.NewGridWrapLayout(size)
}

// fills the container with the correspondent content
func CompleteCards(c *fyne.Container) {
	c.RemoveAll()
	listImagesRecursivelly() // search original images
	getThumbnails()

	// save all images into a go channel to manage concurrently load/generate thumbnails
	channel := make(chan int, len(globals.ImagesList))
	for i := range globals.ImagesList {
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
			card := widget.NewCard("", isolateImageName(globals.ImagesList[i]), content)
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
		wallpaper.SetWallpaper(globals.ImagesList[i])
	}

	// resize the image and get the thumbnails
	resizeImage(i)
	aux := canvas.NewImageFromFile(globals.Thumbnails[i])
	aux.ScaleMode = canvas.ImageScaleFastest
	aux.FillMode = canvas.ImageFillContain

	// A bit of "magia" (With the max layout we can overlap the button and the thumbnail)
	cont := container.NewMax(aux, button)
	return cont
}
