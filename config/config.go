package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/elias-gill/walldo-in-go/wallpaper"
)

type GridSize int

const (
	NORMAL GridSize = iota
	SMALL
	LARGE
)

type Configuration struct {
	WallpfillMode    wallpaper.FillStyle `json:"FillStyle"`
	GridSize         GridSize            `json:"GridSize"`
	WallpaperFolders []string            `json:"Paths"`

	CachePath  string
	ConfigPath string
	ConfigFile string
}

var Config Configuration

func Init() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal("Cannot locate cache directory")
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Cannot locate user config directory")
	}

	conf := Configuration{
		ConfigPath: path.Join(configDir, "walldo"),
		CachePath:  path.Join(cacheDir, "walldo"),
		ConfigFile: path.Join(configDir, "walldo", "config.json"),
		WallpaperFolders: []string{
			path.Join("~", "Pictures"),
			path.Join("~", "Wallpapers"),
			path.Join("~", "Images"),
		},
	}

	err = os.MkdirAll(conf.ConfigPath, 0o770)
	if err != nil {
		log.Fatal("Cannot create config directory " + err.Error())
	}

	err = os.MkdirAll(conf.CachePath, 0o770)
	if err != nil {
		log.Fatal("Cannot create cache directory " + err.Error())
	}

	file, err := os.OpenFile(conf.ConfigFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Cannot open or create config file: %v", err)
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&conf)

	Config = conf
}

func PersistConfig() {
	file, err := os.Create(Config.ConfigFile)
	if err != nil {
		fmt.Printf("could not create the configuration file: %s", err.Error())
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(Config); err != nil {
		fmt.Printf("Could not encode JSON data: %s", err.Error())
	}
}

// Este se queda porque tiene lógica de negocio (procesa los paths)
func GetWallpaperSearchPaths() []string {
	var folders []string
	for _, folder := range Config.WallpaperFolders {
		var err error
		folder, err = expandPath(folder)
		if err != nil {
			log.Print(err)
			continue
		}
		folders = append(folders, folder)
	}
	return folders
}

func expandPath(path string) (string, error) {
	path = os.ExpandEnv(path)

	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		path = strings.Replace(path, "~", homeDir, 1)
	}

	return path, nil
}
