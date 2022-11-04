package globals

import (
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	ImagesList []string
	Thumbnails []string
)

const SYS_OS = runtime.GOOS

// App initializers
var (
	MyApp        = app.NewWithID("walldo")
	Window       = MyApp.NewWindow("Walldo in go")
    WindowHeight = 600
    WindowWidth  = 1020
    // TODO: solve "save window size" on close
	/* WindowHeight = MyApp.Preferences().FloatWithFallback("WindowHeight", 600)
	WindowWidth  = MyApp.Preferences().FloatWithFallback("WindowWidth", 1020) */
)

// Grid config variables
var (
	GridSize   = MyApp.Preferences().StringWithFallback("GridSize", "default")
	GridTitles = MyApp.Preferences().StringWithFallback("GridTitles", "Borderless")
)

// Layout configs
var (
	LayoutStyle  = MyApp.Preferences().StringWithFallback("Layout", "Grid")
	FillStrategy = MyApp.Preferences().StringWithFallback("FillStrategy", "Zoom Fill")
)

// Config files
var (
	ConfigFile     string
	ConfigPath     string
	ThumbnailsPath string
)

// Change config values depending on the OS
// ~/.config/walldo/config.json (unix)
// ~/AppData/Local/walldo/config.json (windows)
func SetGlobalValues() {
	os.Setenv("FYNE_THEME", "dark")
	o, _ := os.UserHomeDir()
	Window.Resize(fyne.NewSize(float32(WindowWidth), float32(WindowHeight)))

	switch SYS_OS {
	case "windows":
		ConfigFile = o + "/AppData/Local/walldo/config.json"
		ConfigPath = o + "/AppData/Local/walldo/"
		ThumbnailsPath = o + "/AppData/Local/walldo/resized_images/"

	default:
		// sistemas Unix (Mac y Linux)
		ConfigFile = o + "/.config/walldo/config.json"
		ConfigPath = o + "/.config/walldo/"
		ThumbnailsPath = o + "/.config/walldo/resized_images/"
	}
}
