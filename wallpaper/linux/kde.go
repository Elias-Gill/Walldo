//go:build linux

package linux

import (
	"os/exec"
	"strconv"

	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

func SetKDE(path string, mode modes.FillStyle) error {
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

func setKDEMode(mode modes.FillStyle) error {
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

func getKDEString(mode modes.FillStyle) string {
	switch mode {
	case modes.FILL_CENTER:
		return "6"
	case modes.FILL_ZOOM:
		return "1"
	case modes.FILL_ORIGINAL:
		return "2"
	case modes.FILL_SCALE:
		return "0"
	case modes.FILL_TILE:
		return "3"
	default:
		panic("invalid walllpaper mode")
	}
}
