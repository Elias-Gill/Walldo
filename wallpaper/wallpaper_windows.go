package wallpaper

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

// Windows registry constants
const (
	spiGetDeskWallpaper = 0x0073
	spiSetDeskWallpaper = 0x0014

	uiParam = 0x0000

	spifUpdateINIFile = 0x01
	spifSendChange    = 0x02
)

// Windows syscalls to modify registry
var (
	user32               = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

// ================================
// Wallpaper engine implementation
// ================================

func AvailableModes() []FillStyle {
	return []FillStyle{
		FILL_CENTER,
		FILL_ORIGINAL,
		FILL_ZOOM,
		FILL_SCALE,
		// FILL_TILE is not natively supported by this Windows registry implementation
	}
}

func SetWallpaper(path string, mode FillStyle) error {
	err := setWallpaperMode(mode)
	if err != nil {
		return err
	}

	// Convert path to windows style paths
	filenameUTF16, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	// Change wallpaper
	systemParametersInfo.Call(
		uintptr(spiSetDeskWallpaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(filenameUTF16)),
		uintptr(spifUpdateINIFile|spifSendChange),
	)

	return nil
}

// ================================
// Helper functions
// ================================

// Modify the registry to set the wallpaper mode before applying the image to the background
func setWallpaperMode(mode FillStyle) error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.SET_VALUE)
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
	case FILL_CENTER:
		style = "0"
	case FILL_ORIGINAL:
		style = "6"
	case FILL_SCALE:
		style = "2"
	default:
		// Use FILL_ZOOM as default to prevent sudden crashes
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_ZOOM")
		style = "22"
	}

	return key.SetStringValue("WallpaperStyle", style)
}
