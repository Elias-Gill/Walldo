// go:build windows

package utils

import (
	"github.com/elias-gill/wallpaper"
)

func SetWallpaper(imageDir string) error {
	wallpaper.SetMode(WallpaperFitMode())
	return wallpaper.SetFromFile(imageDir)
}
