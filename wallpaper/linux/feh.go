package linux

import (
	"errors"
	"os/exec"
)

func setFehBackground(file string, mode string) error {
    return exec.Command("feh", mode, file).Run()
}

func (mode Mode) setFehMode() (string, error) {
	switch mode {
	case Center:
		return "--bg-center", nil
	case Crop:
		return "--bg-scale", nil
	case Fit:
		return "--bg-fill", nil
	case Span:
		return "--bg-max", nil
	case Stretch:
		return "--bg-fill", nil
	case Tile:
		return "--bg-tile", nil
	default:
		return "", errors.New("Invalid mode")
	}
}
