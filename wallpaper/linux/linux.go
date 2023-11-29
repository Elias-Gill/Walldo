//go:build linux || darwin

package linux

import (
	"os"
	"os/exec"
	"strconv"
)

// Desktop contains the current desktop environment on Linux.
// Empty string on all other operating systems.
var Desktop = os.Getenv("XDG_CURRENT_DESKTOP")

// Xserver or Wayland display manager. Used for wayland support
var DisplayServer = os.Getenv("XDG_SESSION_TYPE")

// DesktopSession is used by LXDE on Linux.
var DesktopSession = os.Getenv("DESKTOP_SESSION")

// SetFromFile sets wallpaper from a file path.
// Recibe the mode if feh is the (just one)
func LinuxSetFromFile(file string, mode string) error {
	if isWaylandCompliant() {
		return setForWayland(file, mode)
	}

	if isGNOMECompliant() {
		return setForGnome(file, mode)
	}

	switch Desktop {
	case "KDE":
		return setKDE(file, mode)

	case "X-Cinnamon":
		err := exec.Command("dconf", "write", "/org/cinnamon/desktop/background/picture-options",
			strconv.Quote(getGNOMEString(mode))).Run()
		if err != nil {
			return err
		}
		return exec.Command("dconf", "write", "/org/cinnamon/desktop/background/picture-uri",
			strconv.Quote("file://"+file)).Run()

	case "MATE":
		err := exec.Command("dconf", "write", "/org/mate/desktop/background/picture-options",
			strconv.Quote(getGNOMEString(mode))).Run()
		if err != nil {
			return err
		}
		return exec.Command("dconf", "write", "/org/mate/desktop/background/picture-filename", strconv.Quote(file)).Run()

	case "XFCE":
		return setXFCE(file, mode)

	case "LXDE", "LXQT":
		err := exec.Command("pcmanfm", "--wallpaper-mode", getWaylandString(mode)).Run()
		if err != nil {
			return err
		}
		return exec.Command("pcmanfm", "-w", file).Run()

	case "Deepin":
		err := exec.Command("dconf", "write", "/com/deepin/wrap/gnome/desktop/background/picture-options",
			strconv.Quote(getGNOMEString(mode))).Run()
		if err != nil {
			return err
		}
		return exec.Command("dconf", "write", "/com/deepin/wrap/gnome/desktop/background/picture-uri",
			strconv.Quote("file://"+file)).Run()

	default:
		return setFehBackground(file, mode)
	}
}
