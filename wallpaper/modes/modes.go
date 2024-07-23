package modes

type FillStyle int

const (
	FILL_ZOOM FillStyle = iota
	FILL_CENTER
	FILL_TILE
	FILL_ORIGINAL
	FILL_SCALE
)

const (
	STRING_SCALE    = "Scale"
	STRING_TILE     = "Tile"
	STRING_CENTER   = "Center"
	STRING_ORIGINAL = "Original"
	STRING_ZOOM     = "Zoom fill"
)

var aux = map[string]FillStyle{
	STRING_SCALE:    FILL_SCALE,
	STRING_TILE:     FILL_TILE,
	STRING_CENTER:   FILL_CENTER,
	STRING_ORIGINAL: FILL_ORIGINAL,
	STRING_ZOOM:     FILL_ZOOM,
}

var aux2 = map[FillStyle]string{
	FILL_SCALE:    STRING_SCALE,
	FILL_TILE:     STRING_TILE,
	FILL_CENTER:   STRING_CENTER,
	FILL_ORIGINAL: STRING_ORIGINAL,
	FILL_ZOOM:     STRING_ZOOM,
}

func StrToMode(s string) FillStyle {
	return aux[s]
}

func ModeToStr(s FillStyle) string {
	return aux2[s]
}
