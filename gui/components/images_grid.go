package components

import (
	"log"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

type card struct {
	imgPath   string
	container *fyne.Container
	button    *widget.Button
}

type wallpapersGrid struct {
	container *fyne.Container
	grid      *fyne.Container
}

func NewImageGrid() *wallpapersGrid {
	grid := container.NewWithoutLayout()
	res := &wallpapersGrid{
		grid: grid,
		container: container.New(
			layout.NewPaddedLayout(),
			container.NewScroll(grid)),
	}

	return res
}

func (c wallpapersGrid) GetGridContent() *fyne.Container {
	return c.container
}

func (c *wallpapersGrid) RefreshImgGrid() {
	c.grid.RemoveAll()
	utils.RefreshImagesList()

	cardsChannel := c.generateFrames()

	c.fillContainers(cardsChannel)
}

// fills the image grid with frame containers. Returns a channel with cards that
// are going to be filled latter asynchronously.
func (c wallpapersGrid) generateFrames() chan card {
	// define the cards size
	size := globals.Sizes[globals.GridSize]
	c.grid.Layout = layout.NewGridWrapLayout(fyne.NewSize(size.Width, size.Height))

	imagesList := utils.GetImagesList()

	// Save all images into a go channel to manage concurrently load/generate thumbnails
	// PERF: Addes a new container with a button (without an image) as a empty frame, this makes loading times
	// a lot faster. Then create some go routines to display thumbnails concurrently.
	channel := make(chan card, len(imagesList))
	for _, image := range imagesList {
		channel <- c.newEmptyFrame(image)
	}

	c.grid.Refresh()

	return channel
}

// NOTE: keep this as a separate function
// Creates a new container for the card with a button.
func (c *wallpapersGrid) newEmptyFrame(image string) card {
	button := widget.NewButton("", func() {
		err := wallpaper.SetFromFile(strings.Clone(image))
		if err != nil {
			log.Println(err.Error())
		}
	})
	cont := container.NewMax(button)
	c.grid.Add(cont)

	return card{
		imgPath:   image,
		container: cont,
		button:    button,
	}
}

/*
Recibes the channel with a list of "cards" (image + button inside a container).
generates the thumbnail for the card and refresh the container.
create as many threads as cpus for resizing images to make thumbnails.
*/
func (c wallpapersGrid) fillContainers(channel chan card) {
	log.Println("\n Usando ", runtime.NumCPU()-2, " Hilos")

	for i := 0; i < runtime.NumCPU()-2; i++ {
		go func() {
			for card := range channel {
				// resize the image and get the thumbnail name
				thumbail := utils.ResizeImage(card.imgPath)
				image := canvas.NewImageFromFile(thumbail)
				image.ScaleMode = canvas.ImageScaleFastest
				image.FillMode = canvas.ImageFillContain

				// With the max layout we can overlap the button and the thumbnail
                card.container.Add(card.button)
				card.container.Add(image)
				card.container.Refresh()
			}
		}()
	}
}
