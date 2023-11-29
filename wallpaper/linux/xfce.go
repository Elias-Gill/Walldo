//go:build linux || darwin

package linux

import (
	"os/exec"
	"path"
	"strings"

	"github.com/elias-gill/walldo-in-go/globals"
)

func getXFCEProps(key string) ([]string, error) {
	output, err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--list").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.Trim(string(output), "\n"), "\n")
	var desktops []string

	for _, line := range lines {
		if path.Base(line) == key {
			desktops = append(desktops, line)
		}
	}

	return desktops, nil
}

func getXFCE() (string, error) {
	desktops, err := getXFCEProps("last-image")
	if err != nil || len(desktops) == 0 {
		return "", err
	}

	output, err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", desktops[0]).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func setXFCE(file string, mode string) error {
	err := setXFCEMode(mode)
	if err != nil {
		return err
	}

	desktops, err := getXFCEProps("last-image")
	if err != nil {
		return err
	}
	for _, desktop := range desktops {
		err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", desktop, "--set", file).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func setXFCEMode(mode string) error {
	styles, err := getXFCEProps("image-style")
	if err != nil {
		return err
	}

	for _, style := range styles {
		err = exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", style, "--set", getXFCEString(mode)).Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func getXFCEString(mode string) string {
	switch mode {
	case globals.FILL_CENTER:
		return "1"
	case globals.FILL_ZOOM:
		return "4"
	case globals.FILL_ORIGINAL:
		return "5"
	case globals.FILL_SCALE:
		return "3"
	case globals.FILL_TILE:
		return "2"
	default:
		panic("invalid wallpaper mode")
	}
}
