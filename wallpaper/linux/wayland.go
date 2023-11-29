//go:build linux || darwin
package linux

import (
	"os/exec"
	"strings"

	"github.com/elias-gill/walldo-in-go/globals"
)

func (m Mode) getWaylandString() string {
	switch m {
	case globals.FILL_CENTER:
		return "center"
	case globals.FILL_ORIGINAL:
		return "center"
	case globals.FILL_SCALE:
		return "fit"
	case globals.FILL_ZOOM:
		return "fill"
	case globals.FILL_TILE:
        return "tile"
	default:
		panic("invalid wallpaper mode")
	}
}

// INFO: It depends on swaybg
func setWaylandBackground(file string, mode Mode) error {
	// first kill all instances of swaybg then run swaybg
	exec.Command("killall", "swaybg").Run()

	cmd := exec.Command("swaybg", "-m", mode.getWaylandString(), "-i", file)
	err := cmd.Start()
	if err != nil {
		return err
	}

	// detach the process from walldo
	err = cmd.Process.Release()
	if err != nil {
		return err
	}
	return nil
}

func isWaylandCompliant() bool {
	return strings.Contains(DisplayServer, "wayland")
}
