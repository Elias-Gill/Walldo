package gui

import (
	"log"
	"runtime"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/utils"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

type wallpaperCard struct {
	image     utils.Image
	container *fyne.Container

	// Store a reference to the button for when refreshing the grid
	applyButton *widget.Button
}

type wallpaperGallery struct {
	// Holds the main container that wraps the entire scrollable gallery UI
	rootContainer *fyne.Container

	// Arranges image cards in a fluid grid layout. Separated from rootContainer to
	// enable asynchronously image loading and refreshing.
	gridLayout *fyne.Container

	cards []wallpaperCard
}

func NewGallery() *wallpaperGallery {
	// NOTE: the layout format of "GridWrap" is given in the "RefreshGallery" function because
	// the grid size is defined later, this is just a placeholder.
	grid := container.NewWithoutLayout()
	rootContainer := container.New(
		layout.NewPaddedLayout(),
		container.NewScroll(grid))

	return &wallpaperGallery{
		gridLayout:    grid,
		rootContainer: rootContainer,
	}
}

// Repopulates the wallpaper gallery by clearing existing items,
// creating placeholders, and asynchronously loading images.
func (w *wallpaperGallery) RefreshGallery() {
	w.gridLayout.RemoveAll()
	w.cards = []wallpaperCard{}

	// Generate placeholder cards for images before they are rendered
	for _, img := range utils.ListImages() {
		imgCopy := img.Path
		button := widget.NewButton("", func() {
			err := wallpaper.SetWallpaper(imgCopy)
			if err != nil {
				log.Println(err.Error())
			}
		})
		cont := container.NewStack(button)

		// Construct an empty image card
		card := wallpaperCard{
			image:       img,
			container:   cont,
			applyButton: button,
		}
		w.cards = append(w.cards, card)

		w.gridLayout.Add(card.container)
		w.gridLayout.Layout = layout.NewGridWrapLayout(fyne.NewSize(gridScale()))
	}

	// Refresh the UI to reflect the updated grid layout
	w.gridLayout.Refresh()

	// Start filling the placeholders with images asynchronously
	go w.fillContainers()
}

/*
Generates the thumbnail for the card and refresh the container.

Creates as many threads as ceil(cpus/2), so the app runs faster but the
cpu does not get overwhelmed.
*/
func (c wallpaperGallery) fillContainers() {
	threads := int(runtime.NumCPU() / 2)
	if threads <= 0 {
		threads = 1
	}

	log.Printf("Starting %d threads for image scaling", threads)

	wg := sync.WaitGroup{}

	for k := range c.cards {
		card := c.cards[k]

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			// resize the image and get the thumbnail name
			card.image.GenerateThumbnail()
			image := canvas.NewImageFromFile(card.image.Thumbnail)
			image.ScaleMode = canvas.ImageScaleFastest
			image.FillMode = canvas.ImageFillContain

			// With the max layout we can overlap the button and the thumbnail
			fyne.DoAndWait(func() {
				card.container.Add(card.applyButton)
				card.container.Add(image)
				card.container.Refresh()
			})

			wg.Done()
		}(&wg)

		if k%(threads) == 0 {
			wg.Wait()
		}
	}
}

// Returns the scrollable container that holds the wallpaperGallery
func (w wallpaperGallery) View() *fyne.Container {
	return w.rootContainer
}

// Transform GridSize enums from config to actual pixel size values
// TODO: refactor this and the hole config model
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
