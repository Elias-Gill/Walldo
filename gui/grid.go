package gui

import (
	"log"
	"runtime"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

type card struct {
	image     config.Image
	container *fyne.Container
	button    *widget.Button
}

type wallpapersGrid struct {
	container *fyne.Container
	grid      *fyne.Container
	images    []card
}

func NewImageGrid() *wallpapersGrid {
	grid := container.NewWithoutLayout()

	return &wallpapersGrid{
		grid: grid,
		container: container.New(
			layout.NewPaddedLayout(),
			container.NewScroll(grid)),
	}
}

func (w wallpapersGrid) GetContent() *fyne.Container {
	return w.container
}

func (c *wallpapersGrid) RefreshImgGrid() {
	c.grid.RemoveAll()
	c.images = []card{}

	images := config.ListImages()

	for _, img := range images {
		card := newEmptyFrame(img)

		c.images = append(c.images, card)

		c.grid.Add(card.container)

		c.grid.Layout = layout.NewGridWrapLayout(
			fyne.NewSize(gridScale()),
		)
		c.grid.Refresh()
	}

	go c.fillContainers()
}

// NOTE: keep this as a separate function
// Creates a new empty card.
func newEmptyFrame(image config.Image) card {
	button := widget.NewButton("", func() {
		err := wallpaper.SetWallpaper(strings.Clone(image.Path))
		if err != nil {
			log.Println(err.Error())
		}
	})

	cont := container.NewStack(button)

	return card{
		image:     image,
		container: cont,
		button:    button,
	}
}

/*
Generates the thumbnail for the card and refresh the container.

Creates as many threads as cpus-2, so the app runs faster but the
cpu does not get overwhelmed.
*/
func (c wallpapersGrid) fillContainers() {
	log.Println("\n Usando ", runtime.NumCPU()-2, " Hilos")

	wg := sync.WaitGroup{}

	for k := range c.images {
		card := c.images[k]

		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			// resize the image and get the thumbnail name
			card.image.GenerateThumbnail()
			image := canvas.NewImageFromFile(card.image.Thumbnail)
			image.ScaleMode = canvas.ImageScaleFastest
			image.FillMode = canvas.ImageFillContain

			// With the max layout we can overlap the button and the thumbnail
			card.container.Add(card.button)
			card.container.Add(image)
			card.container.Refresh()

			wg.Done()
		}(&wg)

		// generate only as many thumbnails as number of cpus-2
		if k%(runtime.NumCPU()-2) == 0 {
			wg.Wait()
		}
	}
}

func gridScale() (float32, float32) {
	switch config.GetGridSize() {
	case config.LARGE:
		return 145, 128
	case config.SMALL:
		return 90, 80
	default:
		return 115, 105
	}
}
