package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cognusion/imaging"
	"github.com/elias-gill/walldo-in-go/config"
)

type Image struct {
	Thumbnail string
	Path      string
}

func processPath(root string, processedPaths map[string]struct{}) ([]Image, error) {
	var collectedImages []Image

	walkFunc := func(currentPath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if fileInfo.IsDir() {
			if fileInfo.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		if hasValidExtension(currentPath) {
			collectedImages = append(collectedImages, Image{
				Path:      currentPath,
				Thumbnail: generateUniqName(currentPath),
			})
			return nil
		}

		if _, processed := processedPaths[currentPath]; processed {
			return filepath.SkipDir
		}
		processedPaths[currentPath] = struct{}{}

		// Handle symlinks safely avoiding circular loops
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			resolvedPath, err := filepath.EvalSymlinks(currentPath)
			if err != nil {
				return filepath.SkipDir
			}
			if _, processed := processedPaths[resolvedPath]; processed {
				return nil
			}
			processedPaths[currentPath] = struct{}{}

			symlinkImages, _ := processPath(resolvedPath, processedPaths)
			collectedImages = append(collectedImages, symlinkImages...)
		}

		return nil
	}

	err := filepath.Walk(root, walkFunc)
	return collectedImages, err
}

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

// Optimized check avoiding slice allocations on the heap
func hasValidExtension(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

func (i Image) GenerateThumbnail() error {
	if _, err := os.Stat(i.Thumbnail); err == nil {
		return nil
	}

	src, err := imaging.Open(i.Path)
	if err != nil {
		return err
	}

	// Fast scaling algorithm configuration
	src = imaging.Thumbnail(src, 140, 120, imaging.NearestNeighbor)
	return imaging.Save(src, i.Thumbnail)
}

// Creates an MD5 checksum of the absolute path to prevent thumbnail cross-collisions
func generateUniqName(imagePath string) string {
	hasher := md5.New()
	hasher.Write([]byte(imagePath))
	filename := hex.EncodeToString(hasher.Sum(nil)) + ".jpg"
	return filepath.Join(config.Config.CachePath, filename)
}
