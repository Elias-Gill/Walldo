package linux

import (
	"os/exec"
	"strings"
)

// INFO: It depends on swaybg
func setWaylandBackground(file string, mode string) error {
	// first kill all instances of swaybg
	exec.Command("killall", "swaybg").Run()

	// run swaybg to change the bg
	if mode == "" {
		mode = "fill"
	}

	cmd := exec.Command("swaybg", "-m", mode, "-i", file)
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
