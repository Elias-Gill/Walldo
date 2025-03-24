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

	"github.com/cognusion/imaging"
	"github.com/elias-gill/walldo-in-go/config"
)

type Image struct {
	Thumbnail string
	Path      string
}

// Recursively walks through a directory and collects images found within.
// Returns the collected images instead of modifying a passed slice, making
// the function pure and easier to reason about.
func processPath(root string, processedPaths map[string]struct{}) ([]Image, error) {
	var collectedImages []Image

	walkFunc := func(currentPath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v", currentPath, err)
			return err
		}

		// skip .git folder
		if fileInfo.IsDir() && filepath.Base(currentPath) == ".git" {
			return filepath.SkipDir
		}

		// Append image if this is a file
		if !fileInfo.IsDir() && hasValidExtension(currentPath) {
			collectedImages = append(collectedImages, Image{
				Path:      currentPath,
				Thumbnail: generateUniqName(currentPath),
			})
			return nil
		}

		// Check if the path has already been processed; otherwise, mark it as visited.
		if _, processed := processedPaths[currentPath]; processed {
			log.Printf("Ignoring path (already processed): %s", currentPath)
			return filepath.SkipDir
		}
		processedPaths[currentPath] = struct{}{}

		// Expand symlinks
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			symlinkImages, err := followSymlink(currentPath, processedPaths)
			if err != nil {
				log.Printf("Error following symlink %q: %v", currentPath, err)
				return filepath.SkipDir
			}
			collectedImages = append(collectedImages, symlinkImages...)
			return nil
		}

		return nil
	}

	err := filepath.Walk(root, walkFunc)
	return collectedImages, err
}

// Resolves the given symlink to its actual path and processes its contents while
// preventing symlink cycles and duplicate processing.
func followSymlink(symlink string, processedPaths map[string]struct{}) ([]Image, error) {
	resolvedPath, err := filepath.EvalSymlinks(symlink)
	if err != nil {
		return nil, err
	}

	if _, processed := processedPaths[resolvedPath]; processed {
		log.Printf("Ignoring path (symlink cycle): %s --> %s", symlink, resolvedPath)
		return nil, nil
	}

	processedPaths[symlink] = struct{}{}
	log.Printf("Following symlink: %s --> %s", symlink, resolvedPath)

	return processPath(resolvedPath, processedPaths)
}

// ListImages scans the configured search directories and collects images with valid extensions
// while avoiding duplicate processing and handling symlinks.
func ListImages() []Image {
	processedPaths := make(map[string]struct{})
	var collectedImages []Image

	for _, searchPath := range config.GetWallpaperSearchPaths() {
		images, err := processPath(searchPath, processedPaths)
		if err != nil {
			log.Printf("Error walking folder %q: %v", searchPath, err)
		}
		collectedImages = append(collectedImages, images...)
	}

	return collectedImages
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

func generateUniqName(image string) string {
	h := fnv.New32a()
	name := strings.Split(path.Base(image), ".")[0]
	h.Write([]byte(name))

	filename := strconv.Itoa(int(h.Sum32())) + ".jpg"
	return path.Join(config.GetCachePath(), filename)
}
