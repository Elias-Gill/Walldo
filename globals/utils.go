package globals

import (
	"log"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
)

// Thise are the used config files
// ~/.config/walldo/config.json (unix)
// ~/AppData/Local/walldo/config.json (windows)
func SetupEnvVariables() {
	os.Setenv("FYNE_THEME", "dark")
    Window.Resize(fyne.NewSize(float32(WindowWidth), float32(WindowHeight)))

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot establish users home directory: ", err.Error())
	}

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
