package gui

import (
	"log"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/elias-gill/walldo-in-go/globals"
	global "github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/utils"
	"github.com/elias-gill/walldo-in-go/wallpaper"
)

type wallpapersGrid struct {
    content *fyne.Container
}

// Generates and return a new layout acording to the user configurations
func (c *wallpapersGrid) defineLayout() {
	// default card size
	size := fyne.NewSize(150, 130)
	// other grid sizes
	switch globals.GridSize {
	case "small":
		size = fyne.NewSize(110, 100)
	case "large":
		size = fyne.NewSize(195, 175)
	}
	c.content.Layout = layout.NewGridWrapLayout(size)
}

// fills the container with the correspondent content
func (c wallpapersGrid) fillWithCards() {
	c.content.RemoveAll()
	utils.ListImagesRecursivelly() // search original images
	utils.GetThumbnails()

	// save all images into a go channel to manage concurrently load/generate thumbnails
	channel := make(chan int, len(globals.ImagesList))
	for i := range globals.ImagesList {
		channel <- i
	}

	// create more "threads" to increase performance
	for i := 0; i < runtime.NumCPU()-2; i++ {
		go c.addNewCard(channel)
	}
	print("\n Usando ", runtime.NumCPU()-2, " Hilos")
}

// Recibes the channel with the list of images and creates a new card from the every entry
func (c wallpapersGrid) addNewCard(chanel chan int) {
	for i := range chanel {
		content := generateCardContent(i)

		switch globals.GridTitles {
		// grid without captions
		case "Borderless":
			c.content.Add(content)

		// normal grid with captions
		default:
			card := widget.NewCard("", utils.IsolateImageName(globals.ImagesList[i]), content)
			c.content.Add(card)
		}
		c.content.Refresh()
	}
}

// Creates the new component of the wallpapers grid
// Every component has a container, a button with the position of one image from the images list and its thumbnail
func generateCardContent(i int) *fyne.Container {
	button := widget.NewButton("", nil)
	button.OnTapped = func() {
		// the button has the index of the original image
		err := wallpaper.SetWallpaper(globals.ImagesList[i])
        if err != nil {
            log.Println(err.Error())
        }
	}

	// resize the image and get the thumbnails
	utils.ResizeImage(i)
	aux := canvas.NewImageFromFile(globals.Thumbnails[i])
	aux.ScaleMode = canvas.ImageScaleFastest
	aux.FillMode = canvas.ImageFillContain

	// A bit of "magia" (With the max layout we can overlap the button and the thumbnail)
	cont := container.NewMax(aux, button)
	return cont
}

// template for creating a new button with a custom icon
func newButton(name string, f func(), icon string) *widget.Button {
    if len(icon) > 0 {
        ico := fyne.ThemeIconName(icon)
        return widget.NewButtonWithIcon(name, global.MyApp.Settings().Theme().Icon(ico), f)
    }
    return widget.NewButton(name, f)
}
