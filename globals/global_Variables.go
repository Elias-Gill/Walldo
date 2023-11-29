package globals

import (
	"log"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// TODO: move everything to the utils/config_manager
type Size struct {
	Width  float32
	Height float32
}

// App initializers
var (
	MyApp        = app.NewWithID("walldo")
	Window       = MyApp.NewWindow("Walldo in go")
	WindowHeight = MyApp.Preferences().FloatWithFallback("WindowHeight", 600)
	WindowWidth  = MyApp.Preferences().FloatWithFallback("WindowWidth", 1020)
)

// Grid cards sizes
const SIZE_DEFAULT = "default"
const SIZE_SMALL = "small"
const SIZE_LARGE = "large"

var Sizes map[string]Size = map[string]Size{
	SIZE_SMALL:   {Width: 110, Height: 100},
	SIZE_LARGE:   {Width: 195, Height: 175},
	SIZE_DEFAULT: {Width: 150, Height: 130},
}

var (
	GridSize = MyApp.Preferences().StringWithFallback("GridSize", SIZE_DEFAULT)
)

const FILL_ZOOM = "Zoom Fill"
const FILL_SCALE = "Scale"
const FILL_CENTER = "Center"
const FILL_ORIGINAL = "Original"
const FILL_TILE = "Tile"

var (
	FillStrategy = MyApp.Preferences().StringWithFallback("FillStrategy", FILL_ZOOM)
)

// Config files
var (
	ConfigFile     string
	ConfigPath     string
	ThumbnailsPath string
)

// Thise are the used config files
// ~/.config/walldo/config.json (unix)
// ~/AppData/Local/walldo/config.json (windows)
func SetupEnvVariables() {
	os.Setenv("FYNE_THEME", "dark")
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot establish users home directory: ", err.Error())
	}
	Window.Resize(fyne.NewSize(float32(WindowWidth), float32(WindowHeight)))

	switch runtime.GOOS {
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
