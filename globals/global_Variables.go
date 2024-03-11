package globals

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Size struct {
	Width  float32
	Height float32
}

// App initializers.
var (
	MyApp        = app.NewWithID("walldo")
	Window       = MyApp.NewWindow("Walldo in go")
	WindowHeight = MyApp.Preferences().FloatWithFallback("WindowHeight", 600)
	WindowWidth  = MyApp.Preferences().FloatWithFallback("WindowWidth", 1020)
)

// Grid cards sizes.
const (
	SIZE_DEFAULT = "Default"
	SIZE_SMALL   = "Small"
	SIZE_LARGE   = "Large"
)

var (
	GridSize = MyApp.Preferences().StringWithFallback("GridSize", SIZE_DEFAULT)
)

var Sizes map[string]Size = map[string]Size{
	SIZE_LARGE:   {Width: 195, Height: 175},
	SIZE_DEFAULT: {Width: 115, Height: 105},
	SIZE_SMALL:   {Width: 90, Height: 80},
}

// wallpaper fill strategies.
const (
	FILL_ZOOM     = "Zoom Fill"
	FILL_SCALE    = "Scale"
	FILL_CENTER   = "Center"
	FILL_ORIGINAL = "Original"
	FILL_TILE     = "Tile"
)

var (
	FillStrategy = MyApp.Preferences().StringWithFallback("FillStrategy", FILL_ZOOM)
)

// Config files.
var (
	ConfigFile     string
	ConfigPath     string
	ThumbnailsPath string
)

// Restores the window size of the last time the app has oppened.
func RestoreWindowSize() {
	Window.Resize(fyne.NewSize(float32(WindowWidth), float32(WindowHeight)))
}

// Thise are the used config files.
// ~/.config/walldo/config.json (unix).
// ~/AppData/Local/walldo/config.json (windows).
func InitApp() {
	// set darkmode
	os.Setenv("FYNE_THEME", "dark")

	// set configuration paths
	userConfig, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Cannot establish users home directory: ", err.Error())
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal("Cannot establish users home directory: ", err.Error())
	}

	ConfigPath = userConfig + "walldo/"
	ConfigFile = ConfigPath + "config.json"
	ThumbnailsPath = cacheDir + "/walldo/"

	err = os.MkdirAll(ConfigPath, 0o770)
	if err != nil {
		panic("Cannot create config directory " + err.Error())
	}

	err = os.MkdirAll(ThumbnailsPath, 0o770)
	if err != nil {
		panic("Cannot create cache directory " + err.Error())
	}
}
