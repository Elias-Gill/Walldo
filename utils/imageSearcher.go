package utils

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/elias-gill/walldo-in-go/globals"
)

// Resize the image to create a thumbnail.
// If a thumbnail already exists just do nothing
func resizeImage(i int) {
	destino := globals.ResizedImages[i]
	image := globals.OriginalImages[i]

	// if the thumnail does not exists
	if _, err := os.Stat(destino); err != nil {
		src, _ := imaging.Open(image)
		src = imaging.Thumbnail(src, 200, 150, imaging.Box)
		// save the thumbnail on a folder
		// TODO  make this folder into .cache or /temp
		imaging.Save(src, destino)
	}
}

// TODO  this need a redesign
// Update the resized_images list
func getResizedImages() {
	var res []string
	// set a new entry for the resized_images list with a "unique" name
	for _, image := range globals.OriginalImages {
		dest := globals.ThumbnailsPath + isolateResizedImageName(image) + ".jpg"
		res = append(res, dest) // store the path of the new resized image
	}
	globals.ResizedImages = res // save the result globaly
}

// TODO  I have a good idea for filters here
// TODO  display a dialog error on invalid folders
// Goes trought the configured folders recursivelly and list all the supported image files
func listImagesRecursivelly() {
	// get configured folders from the config file
	globals.OriginalImages = []string{}
	folders := GetConfiguredPaths()

	// loop trought the folder recursivelly
	for _, folder := range folders {
		err := filepath.Walk(folder, func(file string, info os.FileInfo, err error) error {
			if err != nil {
				log.Print(err)
				return err
			}

			// ignore .git files
			if strings.Contains(file, ".git") {
				return filepath.SkipDir
			}
			// ignore directories
			if !info.IsDir() && extensionIsValid(file) {
				globals.OriginalImages = append(globals.OriginalImages, file)
			}
			return nil
		})
		if err != nil {
			log.Print(err)
		}
	}
}

// FUTURE: add more sort styles
// sort images by name
func sortImages(metodo string) {
	if metodo == "default" {
		sort.Strings(globals.OriginalImages)
	}
}

// Determine if the file has a valid extension.
// It can be jpg, jpeg or png.
func extensionIsValid(file string) bool {
	// isolate file extension
	aux := strings.Split(file, ".")
	file = aux[len(aux)-1]

	validos := map[string]int{"jpg": 1, "jpeg": 1, "png": 1}
	_, res := validos[file]
	return res
}

// TODO  change characters size depending of the card size
// Returns the first 12 letters of the name of a image. This is for fitting into the captions
func isolateImageName(name string) string {
	// Change backslashes to normal ones
	name = strings.ReplaceAll(name, `\`, `/`)
	res := strings.Split(name, "/")

	largo := len(res) - 1
	aux := res[largo]
	if len(res[largo]) > 12 {
		aux = res[largo][0:12]
		aux = aux + " ..."
	}
	return aux
}

// Returns a new name for the resized image.
// this name has the format parent+file
func isolateResizedImageName(name string) string {
	name = strings.ReplaceAll(name, `\`, `/`)
	res := strings.Split(name, "/")

	largo := len(res) - 1
	aux := res[largo] + res[largo-1]
	return aux
}
