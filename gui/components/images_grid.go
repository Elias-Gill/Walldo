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

type wallpapersGrid struct {
	content *fyne.Container
}

func NewImageGrid() wallpapersGrid {
	res := wallpapersGrid{content: container.NewWithoutLayout()}
	res.fillGrid()
	return res
}

func (c wallpapersGrid) GetGridContent() *fyne.Container {
	return c.content
}

func (c *wallpapersGrid) RefreshImgGrid() {
	c.content.RemoveAll()
	utils.ListImagesRecursivelly()
	c.fillGrid()
}

type card struct {
	imgPath   string
	container *fyne.Container
	button    *widget.Button
}

// fills the container with the correspondent content
func (c wallpapersGrid) fillGrid() {
	// define the cards size
	size := globals.Sizes[globals.GridSize]
	c.content.Layout = layout.NewGridWrapLayout(fyne.NewSize(size.Width, size.Height))

	imagesList := utils.GetImagesList()

	// Save all images into a go channel to manage concurrently load/generate thumbnails
	// PERF: Addes a new container with a button (without an image) as a empty frame, this makes loading times 
    // a lot faster. Then create some go routines to display thumbnails concurrently.
	channel := make(chan card, len(imagesList))
	for _, image := range imagesList {
		channel <- c.newEmptyFrame(image)
	}
	c.fillContainers(channel)
}

// NOTE: keep this as a separate function
// Creates a new container for the card with a button
func (c *wallpapersGrid) newEmptyFrame(image string) card {
	button := widget.NewButton("", func() {
		err := wallpaper.SetFromFile(strings.Clone(image))
		if err != nil {
			log.Println(err.Error())
		}
	})
	cont := container.NewMax(button)
	c.content.Add(cont)

	return card{
		imgPath:   image,
		container: cont,
		button:    button,
	}
}

// Recibes the channel with a list of "cards" (image + button inside a container).
// generates the thumbnail for the card and refresh the container
func (c wallpapersGrid) fillContainers(channel chan card) {
	// create as many threads as cpus for resizing images to make thumbnails
	print("\n Usando ", runtime.NumCPU()-2, " Hilos")
	for i := 0; i < runtime.NumCPU()-2; i++ {
		go func() {
			for card := range channel {
				// resize the image and get the thumbnail name
				thumbail := utils.ResizeImage(card.imgPath)
				image := canvas.NewImageFromFile(thumbail)
				image.ScaleMode = canvas.ImageScaleFastest
				image.FillMode = canvas.ImageFillContain

				// With the max layout we can overlap the button and the thumbnail
				card.container.RemoveAll()
				card.container.Add(image)
				card.container.Add(card.button)
				card.container.Refresh()
			}
		}()
	}
}
