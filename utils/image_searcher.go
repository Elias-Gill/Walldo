package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/elias-gill/walldo-in-go/globals"
)

// Hace el resize de la imagen y la guarda en el destino
// evita generar un archivo si la imagen ya fue reescalada previamente
func resizeImage(i int) {
	destino := globals.Resized_images[i]
	image := globals.Original_images[i]

	if _, e := os.Stat(destino); e != nil { // si no existe
		src, _ := imaging.Open(image)
		src = imaging.Thumbnail(src, 200, 150, imaging.Box)
		imaging.Save(src, destino)
	}
}

// Actualiza el array "resized_images" con las direcciones de las nuevas imagenes reescaladas
func getResizedImages() {
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
	for _, image := range globals.Original_images {
		destino := path + aislarNombreImagenReescalada(image) + ".jpg"
		res = append(res, destino) // guardar la nueva direccion
	}
	globals.Resized_images = res // guardar la imagenes
}

// Retorna las imagenes recursivamente en las carpetas configuradas
// por el usuario
func listarImagenes() {
	// traer carpetas del archivo de configuracion
	folders := ConfiguredPaths()

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
			if !info.IsDir() && extensionIsValid(file) {
				globals.Original_images = append(globals.Original_images, file)
			}
			return nil
		})

		if err != nil {
			log.Print(err)
		}
	}
}

func ordenarImagenes(metodo string) {
    // TODO  agregar mas metodos de ordenamiento
    if metodo == "default"{
        sort.Strings(globals.Original_images)
    }
}

// comprobar que la extensio del archivo sea la correcta
// Solo jpg, png y jpeg
func extensionIsValid (file string) bool {
    // aislar la extension
    aux := strings.Split(file, ".")
    file = aux[len(aux)-1]

    validos := map[string]int {"jpg": 1, "jpeg": 1, "png": 1}
    _, res := validos[file]
    return res
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
