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
	"github.com/elias-gill/walldo-in-go/fuzzyEngine/matching"
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
	allImages     []utils.Image // Memory cache to avoid heavy disk I/O on every keystroke
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

// RefreshGallery performs a full disk scan and updates the cache
func (w *WallpaperGallery) RefreshGallery() {
	w.allImages = utils.ListImages()
	w.RenderGrid(w.allImages)
}

// FilterGallery processes fuzzy search matching against the in-memory cache
func (w *WallpaperGallery) FilterGallery(query string) {
	if query == "" {
		w.RenderGrid(w.allImages)
		return
	}

	var paths []string
	for _, img := range w.allImages {
		paths = append(paths, img.Path)
	}

	var filtered []utils.Image
	matches := matching.FindAll(query, paths)
	for _, match := range matches {
		filtered = append(filtered, w.allImages[match.Idx])
	}

	w.RenderGrid(filtered)
}

// RenderGrid builds the UI elements for a concrete slice of images
func (w *WallpaperGallery) RenderGrid(images []utils.Image) {
	w.gridLayout.RemoveAll()
	w.cards = make([]wallpaperCard, 0, len(images))

	for _, img := range images {
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

	// Pass an isolated snapshot slice to prevent concurrency race states on rapid typing
	go w.fillContainers(w.cards)
}

func (w *WallpaperGallery) fillContainers(cards []wallpaperCard) {
	workers := runtime.NumCPU() / 2
	if workers <= 0 {
		workers = 1
	}

	jobs := make(chan wallpaperCard, len(cards))
	var wg sync.WaitGroup

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

				fyne.DoAndWait(func() {
					card.container.Add(card.applyButton)
					card.container.Add(img)
					card.container.Refresh()
				})
			}
		}()
	}

	for _, card := range cards {
		jobs <- card
	}
	close(jobs)
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
