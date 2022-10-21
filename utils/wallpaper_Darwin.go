// go:build darwin

package utils

import (
	"github.com/elias-gill/wallpaper"
)

func SetWallpaper(imageDir string) error {
	return wallpaper.SetFromFile(imageDir)
}
