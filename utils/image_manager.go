package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/elias-gill/walldo-in-go/globals"
)

// used for GetImagesList(), so we dont need to re-search for images
var imagesList []string

// Resize the image to create a thumbnail.
// If a thumbnail already exists just do nothing
func ResizeImage(image string) string {
	thumbPath := generateThnPath(image)

	// if the thumnail does not exists
	if _, err := os.Stat(thumbPath); err != nil {
		src, _ := imaging.Open(image)
		src = imaging.Thumbnail(src, 200, 180, imaging.NearestNeighbor)
		// save the thumbnail on a folder
		imaging.Save(src, thumbPath)
	}
	return thumbPath
}

// Goes trought the configured folders recursivelly and list all the supported image files.
func ListImagesRecursivelly() {
	imagesList = []string{}
	folders := GetConfiguredPaths()

	// loop trought folders recursivelly
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
			if !info.IsDir() && hasValidExtension(file) {
				imagesList = append(imagesList, file)
			}
			return nil
		})
		if err != nil {
			log.Print(err)
		}
	}
}

// This returns the image list. The difference from ListImagesRecursivelly is that
// this does not have to search again through the folders in order to improve performance for the
// fuzzy engine
func GetImagesList() []string {
	return imagesList
}

// Returns a new name for an image thumbnail
func generateThnPath(image string) string {
	// replace backslashes with normal slashes (for windows)
	image = strings.ReplaceAll(image, `\`, `/`)
	res := strings.Split(image, "/")
	// generate a thumbnail name with format "parent + file"
	largo := len(res) - 1
	thumbnail := res[largo] + res[largo-1]
	return globals.ThumbnailsPath + thumbnail + ".jpg"
}

// Determine if the file has a valid extension.
// It can be jpg, jpeg or png.
func hasValidExtension(file string) bool {
	// isolate file extension
	aux := strings.Split(file, ".")
	file = aux[len(aux)-1]

	validos := map[string]int{"jpg": 1, "jpeg": 1, "png": 1}
	_, res := validos[file]
	return res
}
