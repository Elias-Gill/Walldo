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

var original_images []string
var resized_images []string

// retorna un tamano dependiendo de la configuracion del usuario
func setGridSize(apl fyne.App) fyne.Size {
	tamano := apl.Preferences().String("layout")
	switch tamano {
	case "small":
		return fyne.NewSize(110, 100)
	case "large":
		return fyne.NewSize(175, 150)
	}
	return fyne.NewSize(150, 130)
}

// rellenar la grilla de imagenes de manera asincrona y utilizando
// concurrencia
func SetGridContent(grid *fyne.Container) {
	listarImagenes()                   // buscar las imagenes
	getResizedImages(&original_images) // images reescaladas

	grid.RemoveAll()
	for i := range original_images {
		rellenar(grid, i)
	}
}

// Crea un elemento para la grilla de imagenes y lo anade a la grilla
// Cada imagen tiene asignado un boton, boton que contiene como texto la posicion de dicha imagen
// en el arreglo de imagenes originales
// Al acabar refresca el contenido
func rellenar(grid *fyne.Container, i int) {
	button := widget.NewButton(strconv.Itoa(i), nil)
	button.OnTapped = func() {
		value, _ := strconv.Atoi(button.Text)
		// el boton contiene el index de la imagen original
		SetWallpaper(original_images[value])
	}
	// imagen rescalada
	resizeImage(i)
	aux := canvas.NewImageFromFile(resized_images[i])
	aux.ScaleMode = canvas.ImageScaleFastest
	aux.FillMode = canvas.ImageFillContain

	// algo de magia (el boton se le superpone a la imagen)
	cont := container.NewMax(aux, button)
	card := widget.NewCard("", aislarNombreImagen(original_images[i]), cont)
	grid.Add(card)
	grid.Refresh()
}

// Retorna la grilla de imagenes a ser mostradas.
func NewContentGrid(apl *fyne.App) (*fyne.Container, *fyne.Container) {
	content_grid := container.New(layout.NewGridWrapLayout(setGridSize(*apl)))
	grid := container.NewScroll(content_grid) // make the grid actually scrollable
	grid.SetMinSize(fyne.NewSize(820, 500))
	return container.NewCenter(grid), content_grid
}

// Actualiza el array "resized_images" con las direcciones de las nuevas imagenes reescaladas
func getResizedImages(original_images *[]string) {
	var res []string
	sys_os := runtime.GOOS
	path, _ := os.UserHomeDir() // home del usuario

	// determinar cual es la carpeta de config dependiendo del OS
	if sys_os == "windows" {
		path += "/AppData/Local/walldo/resized_images/"
	} else { // sistemas Unix (Mac y Linux)
		path += "/.config/walldo/resized_images/"
	}

	// anadir una nueva entrada para la imagen reescalada en el arreglo de nombres
	for _, image := range *original_images {
		destino := path + aislarNombreImagenReescalada(image) + ".jpg"
		res = append(res, destino) // guardar la nueva direccion
	}
	resized_images = res // guardar la imagenes
}

// Hace el resize de la imagen y la guarda en el destino
// evita generar un archivo si la imagen ya fue reescalada previamente
func resizeImage(i int) {
	destino := resized_images[i]
	image := original_images[i]

	if _, e := os.Stat(destino); e != nil { // si no existe
		src, _ := imaging.Open(image)
		src = imaging.Thumbnail(src, 200, 150, imaging.Box)
		imaging.Save(src, destino)
	}
}

// Retorna nombre un nombre para la imagen reescalada
func aislarNombreImagen(name string) string {
	// trasnformar las barras invertidas en windows
	name = strings.ReplaceAll(name, `\`, `/`)
	res := strings.Split(name, "/")

	largo := len(res) - 1
	aux := res[largo]
	if len(res[largo]) > 12 {
		aux = res[largo][0:12]
		aux = aux + " ..."
	}
	return aux // retorna el nombre de la imagen (maximo 12 caracteres)
}

// Retorna un nombre para la imagen reescalada
// retorna el nombre de la imagen como "padrearchivo"
func aislarNombreImagenReescalada(name string) string {
	name = strings.ReplaceAll(name, `\`, `/`)
	res := strings.Split(name, "/")

	largo := len(res) - 1
	aux := res[largo] + res[largo-1]
	return aux
}

// Retorna las imagenes recursivamente en las carpetas configuradas
// por el usuario
func listarImagenes() {
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
	original_images = images
}
