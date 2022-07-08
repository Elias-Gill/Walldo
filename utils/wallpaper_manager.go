package utils

import (

	"github.com/reujab/wallpaper"
)
// const CurrentWallpaper = wallpaper.Get()

func SetWallpaper(imageDir string) error {
    return wallpaper.SetFromFile(imageDir)

}

/* func SetWallpaperFromUrl() {
    err := wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
} */


func WallaperFitMode(mode string) error {
    // TODO  poner este switch con un map (diccionarios)
    switch mode {
    case "Fit":
        return wallpaper.SetMode(wallpaper.Fit)
    case "Crop":
        return wallpaper.SetMode(wallpaper.Crop)
    case "Center":
        return wallpaper.SetMode(wallpaper.Center)
    case "Span":
        return wallpaper.SetMode(wallpaper.Span)
    case "Tile":
        return wallpaper.SetMode(wallpaper.Tile)
    case "Stretch":
        return wallpaper.SetMode(wallpaper.Stretch)
    }
    return nil
}

// retorna el wallpaper actual
func GetCurrentWallpaper() (string, error){
	return wallpaper.Get()
}
