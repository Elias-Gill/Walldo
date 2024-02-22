package utils

import (
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/elias-gill/walldo-in-go/globals"
)

// used for GetImagesList(), so we dont need to re-search for images.
var imagesList []string

// Resize the image to create a thumbnail.
// If a thumbnail already exists just do nothing.
func ResizeImage(image string) string {
	thumbPath := generateThnPath(image)

	// if the thumnail does exists
	if _, err := os.Stat(thumbPath); err == nil {
		return thumbPath
	}

	src, err := imaging.Open(image)
	if err != nil {
		fmt.Println("Image not found: ", image)
		return ""
	}

	src = imaging.Thumbnail(src, 200, 180, imaging.NearestNeighbor)
	imaging.Save(src, thumbPath)

	return thumbPath
}

// Goes trought the configured folders recursivelly and list all the supported image files.
func RefreshImagesList() {
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
// fuzzy engine.
func GetImagesList() []string {
	return imagesList
}

// Returns a new (hashed) path for an image thumbnail.
func generateThnPath(image string) string {
	h := fnv.New32a()
	name := strings.Split(path.Base(image), ".")[0]
	h.Write([]byte(name))

	return globals.ThumbnailsPath + strconv.Itoa(int(h.Sum32())) + ".jpg"
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
