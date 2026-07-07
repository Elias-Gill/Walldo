package components

import "github.com/elias-gill/walldo-in-go/wallpaper"

const (
	stringScale    = "Scale"
	stringTile     = "Tile"
	stringCenter   = "Center"
	stringOriginal = "Original"
	stringZoom     = "Zoom fill"
)

var stringToMode = map[string]wallpaper.FillStyle{
	stringScale:    wallpaper.FILL_SCALE,
	stringTile:     wallpaper.FILL_TILE,
	stringCenter:   wallpaper.FILL_CENTER,
	stringOriginal: wallpaper.FILL_ORIGINAL,
	stringZoom:     wallpaper.FILL_ZOOM,
}

var modeToString = map[wallpaper.FillStyle]string{
	wallpaper.FILL_SCALE:    stringScale,
	wallpaper.FILL_TILE:     stringTile,
	wallpaper.FILL_CENTER:   stringCenter,
	wallpaper.FILL_ORIGINAL: stringOriginal,
	wallpaper.FILL_ZOOM:     stringZoom,
}

func StrToMode(s string) wallpaper.FillStyle {
	return stringToMode[s]
}

func ModeToStr(m wallpaper.FillStyle) string {
	return modeToString[m]
}
