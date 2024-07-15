//go:build linux || darwin

package wallpaper

import (
	"errors"
	"runtime"

	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/wallpaper/linux"
)

var ErrUnsupportedDE = errors.New("Your desktop environment is not supported")

func SetFromFile(file string, mode globals.FillStyle) error {
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
		return []string{
			string(globals.FILL_ZOOM), string(globals.FILL_CENTER),
			string(globals.FILL_TILE), string(globals.FILL_ORIGINAL),
			string(globals.FILL_SCALE)}
	default:
		return []string{string(globals.FILL_ZOOM)}
	}
}
