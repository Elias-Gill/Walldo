package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/disintegration/imaging"
)

// Retorna la grilla de imagenes a ser mostradas.
// Las imagenes se sacan de las carpetas configuradas
// de acuerdo con el archivo de configuracion del usuario
func NewContentGrid() *fyne.Container {
	content_grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(150, 150)))

	original_images := getImages()                       // buscar las imagenes
	resized_images := getResizedImages(&original_images) // images reescaladas

	for i, image := range original_images {
		// boton de accion para la imagen
		button := widget.NewButton(strconv.Itoa(i), nil)
		button.OnTapped = func() {
            value, _:= strconv.Atoi(button.Text)
			SetWallpaper(original_images[value]) // el boton contiene la imagen original
		}

		aux := canvas.NewImageFromFile(resized_images[i]) // imagen rescalada
		aux.ScaleMode = canvas.ImageScaleFastest
		aux.FillMode = canvas.ImageFillContain

		// algo de magia (el boton se le superpone a la imagen)
		cont := container.NewMax(aux, button)
		card := widget.NewCard("", getNombre(image), cont)
		content_grid.Add(card)
	}

	grid := container.NewScroll(content_grid) // make the grid actually scrollable
	grid.SetMinSize(fyne.NewSize(820, 500))
	return container.NewCenter(grid)
}

// crea imagenes reescaladas para mayor rendimiento
// Retorna un array con las direcciones de las nuevas imagenes reescaladas
func getResizedImages(original_images *[]string) []string {
	var res []string
	sys_os := runtime.GOOS
	path, _ := os.UserHomeDir() // home del usuario

	// determinar cual es la carpeta dependiendo del OS
	switch sys_os {
	case "windows":
		path += "/AppData/Local/walldo/resized_images/"
	default:
		// sistemas Unix (Mac y Linux)
		path += "/.config/walldo/resized_images/"
	}

	for _, image := range *original_images {
		// crear la nueva imagen dentro del directorio de reserva
		destino := path + getNombreResize(image) + ".jpg"

		// evita generar un archivo que ya se reescalo previamente
		if _, e := os.Stat(destino); e != nil {
			resizeImage(image, destino)
		}
		res = append(res, destino) // guardar la nueva direccion
	}

	return res
}

// Hace el resize de la imagen y la guarda en el destino
func resizeImage(image string, destino string) {
	src, _ := imaging.Open(image)
	src = imaging.Thumbnail(src, 200, 150, imaging.Box)
	imaging.Save(src, destino)
}

// Retorna nombre un nombre para la imagen reescalada
func getNombre(name string) string {
	name = strings.ReplaceAll(name, `\`, `/`) // trasnformar las barras invertidas en windows(wakala)
	res := strings.Split(name, "/")

	largo := len(res) - 1
	aux := res[largo]
	if len(res[largo]) > 12 {
		aux = res[largo][0:12]
		aux = aux + "..."
	}
	return aux // retorna el nombre de la imagen (maximo 12 caracteres)
}

// Retorna un nombre para la imagen reescalada
// retorna el nombre de la imagen como "padrearchivo"
func getNombreResize(name string) string {
	// trasnformar las barras invertidas en windows(wakala)
	name = strings.ReplaceAll(name, `\`, `/`)
	res := strings.Split(name, "/")

	largo := len(res) - 1
	aux := res[largo] + res[largo-1]
	return aux
}

// Retorna las imagenes recursivamente en las carpetas configuradas
// por el usuario
func getImages() []string {
	// traer carpetas del archivo de configuracion
	folders := ConfiguredPaths()
	var images []string

	// recorrer recursivamente cada una de las carpetas del usuario
	for _, folder := range folders {
		err := filepath.Walk(folder, func(file string, info os.FileInfo, err error) error {
			if err != nil {
				log.Print(err)
				return err
			}

			// ignorar git directories
			if strings.Contains(file, ".git") {
				return filepath.SkipDir
			}
			// ignorar directorios
			if !info.IsDir() {
				images = append(images, file)
			}
			return nil
		})

		if err != nil {
			log.Print(err)
		}
	}

	return images
}
