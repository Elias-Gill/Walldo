//go:build linux || darwin

package wallpaper

import (
	"errors"
	"runtime"

	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/wallpaper/linux"
)

var ErrUnsupportedDE = errors.New("Your desktop environment is not supported")

func SetFromFile(file string) error {
	mode := globals.FillStrategy

	switch runtime.GOOS {
	case "linux":
		return linux.LinuxSetFromFile(file, mode)
	case "darwin":
		return darwinSetFromFile(file)
	default:
		return ErrUnsupportedDE
	}
}

func ListAvailableModes() []string {
	switch runtime.GOOS {
	case "linux":
		return []string{globals.FILL_ZOOM, globals.FILL_CENTER, globals.FILL_TILE, globals.FILL_ORIGINAL, globals.FILL_SCALE}
	default:
		return []string{globals.FILL_ZOOM}
	}
}
