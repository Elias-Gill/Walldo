//go:build linux

package wallpaper

import (
	"os"

	"github.com/elias-gill/walldo-in-go/wallpaper/linux"
	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

// desktop contains the current desktop environment on Linux.
// Empty string on all other operating systems.
var desktop = os.Getenv("XDG_CURRENT_DESKTOP")

// Xserver or Wayland display manager. Used for wayland support
var displayServer = os.Getenv("XDG_SESSION_TYPE")

// session is used by LXDE on Linux.
var session = os.Getenv("DESKTOP_SESSION")

// linux.SetFromFile sets wallpaper from a file path.
// Recibe the mode if feh is the (just one)
func setFromFile(file string, mode modes.FillStyle) error {
	if linux.IsWaylandCompliant(displayServer) {
		return linux.SetWayland(file, mode)
	}

	if linux.IsGNOMECompliant(desktop) {
		return linux.SetGnome(file, mode)
	}

	switch desktop {
	case "KDE":
		return linux.SetKDE(file, mode)
	case "Mate":
		return linux.SetMate(file, mode)
	case "X-Cinnamon":
		return linux.SetKDE(file, mode)
	case "XFCE":
		return linux.SetXFCE(file, mode)
	case "LXDE", "LXQT":
		return linux.SetLxde(file, mode)
	case "Deepin":
		return linux.SetDeepin(file, mode)
	default:
		return linux.SetFehBackground(file, mode)
	}
}
