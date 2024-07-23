//go:build windows

package wallpaper

import (
	"syscall"
	"unsafe"

	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
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
func setFromFile(filename string, mode modes.FillStyle) error {
	err := windowsSetMode(mode)
	if err != nil {
		return err
	}

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
func windowsSetMode(mode modes.FillStyle) error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, "Control Panel\\Desktop", registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.SetStringValue("TileWallpaper", "0")
	if err != nil {
		return err
	}

	var style string
	switch mode {
	case modes.FILL_CENTER:
		style = "0"
	case modes.FILL_ORIGINAL:
		style = "6"
	case modes.FILL_ZOOM:
		style = "22"
	case modes.FILL_SCALE:
		style = "2"
	default:
		panic("invalid wallpaper mode")
	}

	return key.SetStringValue("WallpaperStyle", style)
}
