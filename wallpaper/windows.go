//go:build windows

package wallpaper

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724947.aspx
const (
	spiGetDeskWallpaper = 0x0073
	spiSetDeskWallpaper = 0x0014

	uiParam = 0x0000

	spifUpdateINIFile = 0x01
	spifSendChange    = 0x02
)

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724947.aspx
var (
	user32               = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

// SetFromFile sets the wallpaper for the current user.
func SetFromFile(filename string) error {
	filenameUTF16, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}

	systemParametersInfo.Call(
		uintptr(spiSetDeskWallpaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(filenameUTF16)),
		uintptr(spifUpdateINIFile|spifSendChange),
	)
	return nil
}

// SetMode sets the wallpaper mode.
func windowsSetMode(mode Mode) error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, "Control Panel\\Desktop", registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	var tile string
	if mode == Tile {
		tile = "1"
	} else {
		tile = "0"
	}
	err = key.SetStringValue("TileWallpaper", tile)
	if err != nil {
		return err
	}

	var style string
	switch mode {
	case Center, Tile:
		style = "0"
	case Fit:
		style = "6"
	case Span:
		style = "22"
	case Stretch:
		style = "2"
	case Crop:
		style = "10"
	default:
		panic("invalid wallpaper mode")
	}
	err = key.SetStringValue("WallpaperStyle", style)
	if err != nil {
		return err
	}

	// updates wallpaper
	path, err := Get()
	if err != nil {
		return err
	}

	return windowsSetFromFile(path)
}

func windowsGetCacheDir() (string, error) {
	return os.TempDir(), nil
}
