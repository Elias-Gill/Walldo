package wallpaper

import (
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/wallpaper"
)

func WallpaperFitMode() wallpaper.Mode {
	m := map[string]wallpaper.Mode{
		"Zoom Fill": wallpaper.Fit,
		"Scale":     wallpaper.Crop,
		"Center":    wallpaper.Center,
		"Original":  wallpaper.Span,
		"Tile":      wallpaper.Tile,
	}
	// si es que el map contiene la estrategia
	return m[globals.FillStrategy]
}

// returns the current wallpaper
func GetCurrentWallpaper() (string, error) {
	return wallpaper.Get()
}
