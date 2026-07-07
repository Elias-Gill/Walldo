package components

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/fuzzyEngine/matching"
	"github.com/elias-gill/walldo-in-go/images"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

type wallpaperCard struct {
	image       images.Image
	container   *fyne.Container
	applyButton *widget.Button
}

type WallpaperGallery struct {
	rootContainer *fyne.Container
	gridLayout    *fyne.Container
	cards         []wallpaperCard
	allImages     []images.Image
	allPaths      []string // Cached paths to avoid allocations on every keystroke

	debounceTimer *time.Timer
	cancelWorkers context.CancelFunc // Cancels previous rendering routines
	mu            sync.Mutex
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
	w.mu.Lock()
	w.allImages = images.ListImages()

	// Cache paths immediately to optimize subsequent fuzzy search iterations
	w.allPaths = make([]string, len(w.allImages))
	for i, img := range w.allImages {
		w.allPaths[i] = img.Path
	}
	w.mu.Unlock()

	w.RenderGrid(w.allImages)
}

func (w *WallpaperGallery) FilterGallery(query string) {
	w.mu.Lock()
	if w.debounceTimer != nil {
		w.debounceTimer.Stop()
	}

	w.debounceTimer = time.AfterFunc(110*time.Millisecond, func() {
		if query == "" {
			w.RenderGrid(w.allImages)
			return
		}

		w.mu.Lock()
		cachedPaths := w.allPaths
		cachedImages := w.allImages
		w.mu.Unlock()

		var filtered []images.Image
		matches := matching.FindAll(query, cachedPaths)
		for _, match := range matches {
			filtered = append(filtered, cachedImages[match.Idx])
		}

		w.RenderGrid(filtered)
	})
	w.mu.Unlock()
}

func (w *WallpaperGallery) RenderGrid(targetImages []images.Image) {
	w.mu.Lock()
	// Stop any active thumbnail workers from processing obsolete image batches
	if w.cancelWorkers != nil {
		w.cancelWorkers()
	}
	ctx, cancel := context.WithCancel(context.Background())
	w.cancelWorkers = cancel
	w.mu.Unlock()

	// Prepare data state off the main UI thread
	cards := make([]wallpaperCard, len(targetImages))
	for i, img := range targetImages {
		imgCopy := img.Path
		button := widget.NewButton("", func() {
			err := wallpaper.SetWallpaper(imgCopy, config.Config.WallpfillMode)
			if err != nil {
				log.Println("Error setting wallpaper:", err)
			}
		})
		cards[i] = wallpaperCard{
			image:       img,
			container:   container.NewStack(button),
			applyButton: button,
		}
	}

	// Safely perform structural mutations inside the main GUI thread loop
	fyne.DoAndWait(func() {
		w.gridLayout.RemoveAll()
		w.cards = cards
		for _, card := range cards {
			w.gridLayout.Add(card.container)
		}
		w.gridLayout.Layout = layout.NewGridWrapLayout(fyne.NewSize(gridScale()))
		w.gridLayout.Refresh()
	})

	go w.fillContainers(ctx, cards)
}

func (w *WallpaperGallery) fillContainers(ctx context.Context, cards []wallpaperCard) {
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
				select {
				case <-ctx.Done():
					return
				default:
				}

				if err := card.image.GenerateThumbnail(); err != nil {
					continue
				}

				img := canvas.NewImageFromFile(card.image.Thumbnail)
				img.ScaleMode = canvas.ImageScaleFastest
				img.FillMode = canvas.ImageFillContain

				select {
				case <-ctx.Done():
					return
				default:
				}

				fyne.DoAndWait(func() {
					card.container.Add(card.applyButton)
					card.container.Add(img)
					card.container.Refresh()
				})
			}
		}()
	}

	for _, card := range cards {
		select {
		case <-ctx.Done():
			close(jobs)
			return
		default:
			jobs <- card
		}
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
