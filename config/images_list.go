package config

import (
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cognusion/imaging"
)

type Image struct {
	Thumbnail string
	Path      string
}

func ListImages() []Image {
	return conf.searchImages()
}

// Goes trought the configured folders recursivelly and list all the supported image files.
func (c Configuration) searchImages() []Image {
	imagesList := []Image{}

	// loop trought folders recursivelly
	for _, folder := range c.Paths {
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
				imagesList = append(imagesList, Image{
					Path: file,
					Thumbnail: func(image string) string {
						h := fnv.New32a()
						name := strings.Split(path.Base(image), ".")[0]
						h.Write([]byte(name))

						return c.cachePath + strconv.Itoa(int(h.Sum32())) + ".jpg"
					}(file),
				})
			}

			return nil
		})
		if err != nil {
			log.Print(err)
		}
	}

	return imagesList
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

// Lazyly generate thumbnails
// If a thumbnail already exists just do nothing.
func (i Image) GenerateThumbnail() error {
	// if the thumbnail does exists
	if _, err := os.Stat(i.Thumbnail); err == nil {
		return nil
	}

	src, err := imaging.Open(i.Path)
	if err != nil {
		fmt.Println("Image not found: ", i.Path)
		return err
	}

	src = imaging.Thumbnail(src, 140, 120, imaging.NearestNeighbor)
	imaging.Save(src, i.Thumbnail)

	return nil
}
