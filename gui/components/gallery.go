package components

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
	image       utils.Image
	container   *fyne.Container
	applyButton *widget.Button
}

type WallpaperGallery struct {
	rootContainer *fyne.Container
	gridLayout    *fyne.Container
	cards         []wallpaperCard
}

func NewGallery() *WallpaperGallery {
	grid := container.NewWithoutLayout()
	rootContainer := container.New(
		layout.NewPaddedLayout(),
		container.NewScroll(grid),
	)

	return &WallpaperGallery{
		gridLayout:    grid,
		rootContainer: rootContainer,
	}
}

func (w *WallpaperGallery) View() *fyne.Container {
	return w.rootContainer
}

func (w *WallpaperGallery) RefreshGallery() {
	w.gridLayout.RemoveAll()
	w.cards = []wallpaperCard{}

	// Draw UI placeholders immediately to preserve UI responsiveness
	for _, img := range utils.ListImages() {
		imgCopy := img.Path
		button := widget.NewButton("", func() {
			err := wallpaper.SetWallpaper(imgCopy, config.Config.WallpfillMode)
			if err != nil {
				log.Println("Error setting wallpaper:", err)
			}
		})
		cont := container.NewStack(button)

		card := wallpaperCard{
			image:       img,
			container:   cont,
			applyButton: button,
		}
		w.cards = append(w.cards, card)
		w.gridLayout.Add(card.container)
	}

	w.gridLayout.Layout = layout.NewGridWrapLayout(fyne.NewSize(gridScale()))
	w.gridLayout.Refresh()

	// Process images using a decoupled worker pool
	go w.fillContainers()
}

func (w *WallpaperGallery) fillContainers() {
	workers := runtime.NumCPU() / 2
	if workers <= 0 {
		workers = 1
	}

	// Channel to pipe scaling jobs to the worker pool
	jobs := make(chan wallpaperCard, len(w.cards))
	var wg sync.WaitGroup

	// Start a fixed number of persistent worker routines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for card := range jobs {
				if err := card.image.GenerateThumbnail(); err != nil {
					continue
				}

				img := canvas.NewImageFromFile(card.image.Thumbnail)
				img.ScaleMode = canvas.ImageScaleFastest
				img.FillMode = canvas.ImageFillContain

				// Thread-safe UI injection via Fyne runtime event loop
				fyne.DoAndWait(func() {
					card.container.Add(card.applyButton)
					card.container.Add(img)
					card.container.Refresh()
				})
			}
		}()
	}

	// Feed all cards into the concurrent processing queue
	for _, card := range w.cards {
		jobs <- card
	}
	close(jobs)

	// Keep workers alive until the queue is completely drained
	wg.Wait()
}

func gridScale() (float32, float32) {
	switch config.Config.GridSize {
	case config.LARGE:
		return 145, 128
	case config.SMALL:
		return 90, 80
	default:
		return 115, 105
	}
}
