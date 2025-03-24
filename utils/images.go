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

var recursiveFolders = make(map[string]struct{})

func followSymlink(symlink string) ([]Image, error) {
    realPath, err := filepath.EvalSymlinks(symlink)
    if err != nil {
        return nil, err
    }

    // Only check if we've already processed this exact real path
    if _, visited := recursiveFolders[realPath]; visited {
        log.Printf("Ignoring path (symlink cycle): %s", symlink)
        return nil, nil
    }

    // Mark as visited after we've checked
    recursiveFolders[realPath] = struct{}{}
    log.Printf("Following symlink: %s", symlink)

    var imagesList []Image

    err = filepath.Walk(realPath, func(dirPath string, dirInfo os.FileInfo, err error) error {
        if err != nil {
            log.Printf("Error accessing path %q: %v", dirPath, err)
            return err
        }

        // Only check if we've already processed this exact real path
        if _, visited := recursiveFolders[dirPath]; visited {
            log.Printf("Ignoring path (already processed): %s", dirPath)
            return filepath.SkipDir
        }

        // Mark as visited after we've checked
        recursiveFolders[dirPath] = struct{}{}

        // Ignore .git directory
        if dirInfo.IsDir() && filepath.Base(dirPath) == ".git" {
            return filepath.SkipDir
        }

        // Handle symlinks
        if dirInfo.Mode()&os.ModeSymlink != 0 {
            list, err := followSymlink(dirPath)
            if err != nil {
                log.Printf("Error following symlink %q: %v", dirPath, err)
                return nil // Skip this symlink but continue walking
            }
            imagesList = append(imagesList, list...)
            return nil
        }

        // Process files with valid extensions
        if !dirInfo.IsDir() && hasValidExtension(dirPath) {
            imagesList = append(imagesList, Image{
                Path:      dirPath,
                Thumbnail: generateUniqName(dirPath),
            })
        }

        return nil
    })

    return imagesList, err
}

func ListImages() []Image {
	var imagesList []Image

	// Clear the map before starting
	recursiveFolders = make(map[string]struct{})

	for _, folder := range config.GetWallpaperSearchPaths() {
		err := filepath.Walk(folder, func(dirPath string, dirInfo os.FileInfo, err error) error {
			if err != nil {
				log.Printf("Error accessing path %q: %v", dirPath, err)
				return err
			}

			// Ignore already visited folders
			if _, visited := recursiveFolders[dirPath]; visited {
				log.Printf("Ignoring path (already visited): %s", dirPath)
				return filepath.SkipDir
			}
			// Mark path as visited
			recursiveFolders[dirPath] = struct{}{}

			// Ignore .git directory
			if dirInfo.IsDir() && filepath.Base(dirPath) == ".git" {
				return filepath.SkipDir
			}

			// Handle symlinks
			if dirInfo.Mode()&os.ModeSymlink != 0 {
				list, err := followSymlink(dirPath)
				if err != nil {
					log.Printf("Error following symlink %q: %v", dirPath, err)
					return nil
				}
				imagesList = append(imagesList, list...)
				return nil
			}

			// Process files with valid extensions
			if !dirInfo.IsDir() && hasValidExtension(dirPath) {
				imagesList = append(imagesList, Image{
					Path:      dirPath,
					Thumbnail: generateUniqName(dirPath),
				})
			}

			return nil
		})

		if err != nil {
			log.Printf("Error walking folder %q: %v", folder, err)
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

func generateUniqName(image string) string {
	h := fnv.New32a()
	name := strings.Split(path.Base(image), ".")[0]
	h.Write([]byte(name))

	filename := strconv.Itoa(int(h.Sum32())) + ".jpg"
	return path.Join(config.GetCachePath(), filename)
}
