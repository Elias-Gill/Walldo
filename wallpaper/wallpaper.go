//go:build linux || darwin

package wallpaper

import (
	"errors"
	"runtime"

	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/wallpaper/linux"
)

type Mode int

const (
	Center Mode = iota
	Crop
	Fit
	Span
	Stretch
	Tile
)

// ErrUnsupportedDE is thrown when Desktop is not a supported desktop environment.
var ErrUnsupportedDE = errors.New("your desktop environment is not supported")

func SetFromFile(file string) error {
	switch runtime.GOOS {
	case "linux":
        // FIX: serious problems with this shitty code I wrote
        // fix the problems with changing image modes
		return linux.LinuxSetFromFile(file)
	case "darwin":
		return darwinSetFromFile(file)
	}
	return ErrUnsupportedDE
}

func WallpaperFitMode() Mode {
	switch globals.FillStrategy {
	case globals.FILL_ZOOM:
		return Fit
	case globals.FILL_SCALE:
		return Crop
	case globals.FILL_CENTER:
		return Center
	case globals.FILL_ORIGINAL:
		return Span
	case globals.FILL_TILE:
		return Tile
	}
	return Fit
}
