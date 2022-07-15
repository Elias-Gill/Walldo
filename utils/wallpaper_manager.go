package utils

import (
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/wallpaper"
)

// const CurrentWallpaper = wallpaper.Get()

func SetWallpaper(imageDir string) error {
    mode, _ := wallpaper.SetMode(WallaperFitMode())
	return wallpaper.SetFromFile(imageDir, mode)
}

/* func SetWallpaperFromUrl() {
    err := wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
} */

func WallaperFitMode() wallpaper.Mode {
	// TODO  poner este switch con un map (diccionarios)
	switch globals.FillStrategy {
	case "Fit":
		return wallpaper.Fit
	case "Crop":
		return wallpaper.Crop
	case "Center":
		return wallpaper.Center
	case "Span":
		return wallpaper.Span
	case "Tile":
		return wallpaper.Tile
	case "Stretch":
		return wallpaper.Stretch
	}
    return wallpaper.Fit
}

// retorna el wallpaper actual
func GetCurrentWallpaper() (string, error) {
	return wallpaper.Get()
}
