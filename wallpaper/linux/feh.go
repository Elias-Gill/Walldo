//go:build linux

package linux

import (
	"os/exec"

	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

func SetFehBackground(file string, mode modes.FillStyle) error {
	return exec.Command("feh", getFehString(mode), file).Run()
}

func getFehString(mode modes.FillStyle) string {
	switch mode {
	case modes.FILL_CENTER:
		return "--bg-center"
	case modes.FILL_SCALE:
		return "--bg-scale"
	case modes.FILL_ZOOM:
		return "--bg-fill"
	case modes.FILL_ORIGINAL:
		return "--bg-max"
	case modes.FILL_TILE:
		return "--bg-tile"
	default:
		panic("invalid wallpaper mode")
	}
}
