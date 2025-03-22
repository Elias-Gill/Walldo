package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/gui"
)

func main() {
	uninstallFlag := flag.Bool("uninstall", false, "Uninstall the executable by deleting itself")
	flag.Parse()

	if *uninstallFlag {
		// Get the path to the current executable
		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Error getting executable path:", err)
			return
		}

		// Attempt to delete the executable
		err = os.Remove(exePath)
		if err != nil {
			fmt.Println("Error deleting executable:", err)
			return
		}

		fmt.Println("Executable uninstalled successfully:", exePath)
		return
	}

	startGui()
}

func startGui() {
	app := app.NewWithID("Walldo")
	window := app.NewWindow("Walldo")

	config.InitConfig(window, app.Settings())

	// title
	title := canvas.NewText("Select your wallpaper", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 18

	// scrollable image grid
	grid := gui.NewGallery()

	// NOTE: NewBottomNav receives a callback to refresh the gallery.
	// While this approach may not be the most elegant, introducing a complex
	// event manager would be overkill for the scope of this project.
	nav := gui.NewBottomNav(grid.RefreshGallery)

	window.SetContent(
		container.New(
			layout.NewBorderLayout(title, nav, nil, nil),
			title,
			grid.View(),
			nav,
		),
	)

	// load images and thumbnails while initializing the GUI
	go grid.RefreshGallery()

	// save the window size on close
	app.Lifecycle().SetOnStopped(func() {
		app.Preferences().SetFloat("WindowHeight", float64(window.Canvas().Size().Height))
		app.Preferences().SetFloat("WindowWidth", float64(window.Canvas().Size().Width))
		config.PersistConfig()
	})

	// restore previous window size
	window.Resize(
		fyne.NewSize(
			float32(app.Preferences().FloatWithFallback("WindowWidth", 800)),
			float32(app.Preferences().FloatWithFallback("WindowHeight", 800)),
		),
	)

	window.ShowAndRun()
}
