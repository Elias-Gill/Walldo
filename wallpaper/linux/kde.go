//go:build linux || darwin

package linux

import (
	"os/exec"
	"strconv"

	"github.com/elias-gill/walldo-in-go/globals"
)

func setKDE(path string, mode globals.FillStyle) error {
	err := setKDEMode(mode)
	if err != nil {
		return err
	}

	return evalKDE(`
		for (const desktop of desktops()) {
			desktop.currentConfigGroup = ["Wallpaper", "org.kde.image", "General"]
			desktop.writeConfig("Image", ` + strconv.Quote("file://"+path) + `)
		}
	`)
}

func setKDEMode(mode globals.FillStyle) error {
	return evalKDE(`
		for (const desktop of desktops()) {
			desktop.currentConfigGroup = ["Wallpaper", "org.kde.image", "General"]
			desktop.writeConfig("FillMode", ` + getKDEString(mode) + `)
		}
	`)
}

func evalKDE(script string) error {
	return exec.Command("qdbus", "org.kde.plasmashell", "/PlasmaShell", "org.kde.PlasmaShell.evaluateScript", script).Run()
}

func getKDEString(mode globals.FillStyle) string {
	switch mode {
	case globals.FILL_CENTER:
		return "6"
	case globals.FILL_ZOOM:
		return "1"
	case globals.FILL_ORIGINAL:
		return "2"
	case globals.FILL_SCALE:
		return "0"
	case globals.FILL_TILE:
		return "3"
	default:
		panic("invalid walllpaper mode")
	}
}
