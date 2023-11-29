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
	}
	return ErrUnsupportedDE
}
