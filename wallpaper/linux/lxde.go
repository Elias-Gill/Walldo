//go:build linux
package linux

import (
	"os/exec"

	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

func SetLxde(file string, mode modes.FillStyle) error {
	err := exec.Command("pcmanfm", "--wallpaper-mode", getWaylandString(mode)).Run()
	if err != nil {
		return err
	}
	return exec.Command("pcmanfm", "-w", file).Run()
}
