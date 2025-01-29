package wallpaper

import (
	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

var mode modes.FillStyle = modes.FILL_ORIGINAL

func SetWallpaper(file string) error {
	return setFromFile(file, mode)
}

func GetCurMode() modes.FillStyle {
	return mode
}

func SetMode(m modes.FillStyle) {
	mode = m
}

func ListModes() []string {
	return []string{
		modes.STRING_SCALE,
		modes.STRING_TILE,
		modes.STRING_CENTER,
		modes.STRING_ORIGINAL,
		modes.STRING_ZOOM,
	}
}
