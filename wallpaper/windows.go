//go:build windows

package wallpaper

import (
	"os"
	"syscall"
	"unsafe"

	"github.com/elias-gill/walldo-in-go/globals"
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
	err := windowsSetMode()
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
func windowsSetMode() error {
	mode := globals.FillStrategy

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
	case globals.FILL_CENTER:
		style = "0"
	case globals.FILL_ORIGINAL:
		style = "6"
	case globals.FILL_ZOOM:
		style = "22"
	case globals.FILL_SCALE:
		style = "2"
	default:
		panic("invalid wallpaper mode")
	}

	return key.SetStringValue("WallpaperStyle", style)
}

func ListAvailableModes() []string {
	return []string{
		globals.FILL_ZOOM, globals.FILL_CENTER,
		globals.FILL_ORIGINAL, globals.FILL_SCALE}
}

func windowsGetCacheDir() (string, error) {
	return os.TempDir(), nil
}
