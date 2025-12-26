//go:build linux

package linux

import (
	"os/exec"
	"os"
	"strings"

	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

func getWaylandString(mode modes.FillStyle) string {
	switch mode {
	case modes.FILL_CENTER:
		return "center"
	case modes.FILL_ORIGINAL:
		return "center"
	case modes.FILL_SCALE:
		return "fit"
	case modes.FILL_ZOOM:
		return "fill"
	case modes.FILL_TILE:
		return "tile"
	default:
		panic("invalid wallpaper mode")
	}
}

// INFO: It depends on swaybg
func SetWayland(file string, mode modes.FillStyle) error {
	// first kill all instances of swaybg
	exec.Command("killall", "swaybg").Run()

	// save the current wallpaper file to $HOME/.config/waybackground
	configPath := os.ExpandEnv("$HOME/.config/swaybg")
	if err := os.WriteFile(configPath, []byte(file+"\n"), 0644); err != nil {
		return err
	}

	cmd := exec.Command("swaybg", "-m", getWaylandString(mode), "-i", file)
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

func IsWaylandCompliant(displayServer string) bool {
	return strings.Contains(displayServer, "wayland")
}
