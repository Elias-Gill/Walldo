//go:build linux || darwin

package wallpaper

import (
	"errors"
	"runtime"

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
		return linux.LinuxSetFromFile(file)
	case "darwin":
		return darwinSetFromFile(file)
	}
    return ErrUnsupportedDE
}
