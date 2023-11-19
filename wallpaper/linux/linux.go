package linux

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
)

// Desktop contains the current desktop environment on Linux.
// Empty string on all other operating systems.
var Desktop = os.Getenv("XDG_CURRENT_DESKTOP")

// Xserver or Wayland display manager. Used for wayland support
var DisplayServer = os.Getenv("XDG_SESSION_TYPE")

// DesktopSession is used by LXDE on Linux.
var DesktopSession = os.Getenv("DESKTOP_SESSION")

type Mode int

const (
	Center Mode = iota
	Crop
	Fit
	Span
	Stretch
	Tile
)

// SetFromFile sets wallpaper from a file path.
// Recibe the mode if feh is the (just one)
func LinuxSetFromFile(file string, mode ...string) error {
	if isWaylandCompliant() {
		return setWaylandBackground(file, mode[0])
	}

	if isGNOMECompliant() {
		return setForGnome(file)
	}

	switch Desktop {
	case "KDE":
		return setKDE(file)
	case "X-Cinnamon":
		return exec.Command("dconf", "write", "/org/cinnamon/desktop/background/picture-uri", strconv.Quote("file://"+file)).Run()
	case "MATE":
		return exec.Command("dconf", "write", "/org/mate/desktop/background/picture-filename", strconv.Quote(file)).Run()
	case "XFCE":
		return setXFCE(file)
	case "LXDE":
		return exec.Command("pcmanfm", "-w", file).Run()
	case "Deepin":
		return exec.Command("dconf", "write", "/com/deepin/wrap/gnome/desktop/background/picture-uri", strconv.Quote("file://"+file)).Run()
	default:
		if len(mode) > 0 {
			return setFehBackground(file, mode[0])
		}
		return setFehBackground(file, "--bg-fill")
	}
}

// SetMode sets the wallpaper mode.
// In case of non DE's this returns a string containing a feh command mode
// You can pass this string to "SetFromFile" to change the feh mode
func linuxSetMode(mode Mode) (string, error) {
	if isGNOMECompliant() {
		return "", exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-options", strconv.Quote(mode.getGNOMEString())).Run()
	}

	switch Desktop {
	case "KDE":
		return "", setKDEMode(mode)
	case "X-Cinnamon":
		return "", exec.Command("dconf", "write", "/org/cinnamon/desktop/background/picture-options", strconv.Quote(mode.getGNOMEString())).Run()
	case "MATE":
		return "", exec.Command("dconf", "write", "/org/mate/desktop/background/picture-options", strconv.Quote(mode.getGNOMEString())).Run()
	case "XFCE":
		return "", setXFCEMode(mode)
	case "LXDE":
		return "", exec.Command("pcmanfm", "--wallpaper-mode", mode.getLXDEString()).Run()
	case "LXQT":
		return "", exec.Command("pcmanfm", "--wallpaper-mode", mode.getLXDEString()).Run()
	case "Deepin":
		return "", exec.Command("dconf", "write", "/com/deepin/wrap/gnome/desktop/background/picture-options", strconv.Quote(mode.getGNOMEString())).Run()
	default:
		return mode.setFehMode()
	}
}

func linuxGetCacheDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, ".cache"), nil
}
