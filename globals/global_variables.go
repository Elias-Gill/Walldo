package globals

import (
	"os"
	"runtime"

	"fyne.io/fyne/v2/app"
)

var (
	OriginalImages []string
	Original_paths  []string
	ResizedImages  []string
)

const SYS_OS = runtime.GOOS

// App initializers
var (
	MyApp  = app.NewWithID("walldo")
	Window = MyApp.NewWindow("Walldo in go")
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
	ConfigDir, _ = os.UserHomeDir() // Home folder
	ConfigPath   = ConfigDir        // configs path
)

// Change config values depending on the OS
// ~/.config/walldo/config.json (unix)
// ~/AppData/Local/walldo/config.json (windows)
func SetGlobalValues() {
	os.Setenv("FYNE_THEME", "dark")

	switch SYS_OS {
	case "windows":
		ConfigDir += "/AppData/Local/walldo/config.json"
		ConfigPath += "/AppData/Local/walldo/"

	default:
		// sistemas Unix (Mac y Linux)
		ConfigDir += "/.config/walldo/config.json"
		ConfigPath += "/.config/walldo/"
	}
}
