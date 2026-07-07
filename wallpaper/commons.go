package wallpaper

// Common interfaces and logic that remains equal on every engine implementation

type FillStyle int

const (
	FILL_ZOOM FillStyle = iota
	FILL_CENTER
	FILL_TILE
	FILL_ORIGINAL
	FILL_SCALE
)
