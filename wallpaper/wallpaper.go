package wallpaper

import (
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/wallpaper"
)

func WallpaperFitMode() wallpaper.Mode {
	switch globals.FillStrategy {
	case "Zoom Fill":
		return wallpaper.Fit
	case "Scale":
		return wallpaper.Crop
	case "Center":
		return wallpaper.Center
	case "Original":
		return wallpaper.Span
	case "Tile":
		return wallpaper.Tile
	}
	return wallpaper.Fit
}

// returns the current wallpaper
func GetCurrentWallpaper() (string, error) {
	return wallpaper.Get()
}

func SetWallpaper(imageDir string) error {
    mode, _ := wallpaper.SetMode(WallpaperFitMode())
    return wallpaper.SetFromFile(imageDir, mode)
}
