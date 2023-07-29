package globals

import (
	"log"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	ImagesList []string
)

const SYS_OS = runtime.GOOS

// App initializers
var (
	MyApp        = app.NewWithID("walldo")
	Window       = MyApp.NewWindow("Walldo in go")
	WindowHeight = MyApp.Preferences().FloatWithFallback("WindowHeight", 600)
	WindowWidth  = MyApp.Preferences().FloatWithFallback("WindowWidth", 1020)
)

// Grid config variables
var (
	GridSize   = MyApp.Preferences().StringWithFallback("GridSize", "default")
	GridTitles = MyApp.Preferences().StringWithFallback("GridTitles", "Borderless")
)

// Layout styles
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
func SetupEnvVariables() {
	os.Setenv("FYNE_THEME", "dark")
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot establish users home directory: ", err.Error())
	}
	Window.Resize(fyne.NewSize(float32(WindowWidth), float32(WindowHeight)))

	switch SYS_OS {
	case "windows":
		ConfigPath = home + "/AppData/Local/walldo/"
		ConfigFile = home + "/AppData/Local/walldo/config.json"
		ThumbnailsPath = home + "/AppData/Local/walldo/resized_images/"

	default:
		// sistemas Unix (Mac y Linux)
		ConfigPath = home + "/.config/walldo/"
		ConfigFile = home + "/.config/walldo/config.json"
		ThumbnailsPath = home + "/.config/walldo/resized_images/"
	}
}
