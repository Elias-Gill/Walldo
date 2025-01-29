package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"fyne.io/fyne/v2"
	"github.com/elias-gill/walldo-in-go/wallpaper"
	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

type GridSize int

const (
	NORMAL GridSize = iota
	SMALL
	LARGE
)

type Configuration struct {
	WallpfillMode    modes.FillStyle `json:"FillStyle"`
	GridSize         GridSize        `json:"GridSize"`
	WallpaperFolders []string        `json:"Paths"`

	cachePath  string
	configPath string
	configFile string

	fyneSettings fyne.Settings
	window       fyne.Window
}

var Config Configuration

func InitConfig(window fyne.Window, fyneSettings fyne.Settings) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal("Cannot locate cache directory")
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Cannot locate user config directory")
	}

	conf := Configuration{
		configPath: path.Join(configDir, "walldo"),
		cachePath:  path.Join(cacheDir, "walldo"),
		configFile: path.Join(configDir, "walldo", "config.json"),
	}

	err = os.MkdirAll(conf.configPath, 0o770)
	if err != nil {
		log.Fatal("Cannot create config directory " + err.Error())
	}

	err = os.MkdirAll(conf.cachePath, 0o770)
	if err != nil {
		log.Fatal("Cannot create cache directory " + err.Error())
	}

	// Open the file if it exists, or create it if it does not exist
	file, err := os.OpenFile(conf.configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Cannot open or create config file: %v", err)
	}
	defer file.Close()

	// Merge users config to default
	json.NewDecoder(file).Decode(&conf)

	// Update global config
	Config = conf
	Config.fyneSettings = fyneSettings
	Config.window = window
}

func PersistConfig() {
	// Open the file (create if it doesn't exist, truncate if it does)
	file, err := os.Create(Config.configFile)
	if err != nil {
		fmt.Printf("could not create the configuration file: %s", err.Error())
		return
	}
	defer file.Close()

	// Create a JSON encoder and set it to write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print with indentation

	// Encode the struct to JSON and write it to the file
	if err := encoder.Encode(Config); err != nil {
		fmt.Printf("Could not encode JSON data: %s", err.Error())
	}
}

func SetGridSize(s GridSize) {
	Config.GridSize = s
}

func GetGridSize() GridSize {
	return Config.GridSize
}

func GetWallpaperSearchPaths() []string {
	return Config.WallpaperFolders
}

func SetPaths(s []string) {
	Config.WallpaperFolders = s
}

func SetWallpFillMode(m modes.FillStyle) {
	Config.WallpfillMode = m
	wallpaper.SetMode(m)
}

func GetWallpFillMode() modes.FillStyle {
	return Config.WallpfillMode
}

func GetFyneSettings() fyne.Settings {
	return Config.fyneSettings
}

func SetFyneSettings(s fyne.Settings) {
	Config.fyneSettings = s
}

func GetWindow() fyne.Window {
	return Config.window
}

func GetCachePath() string {
	return Config.cachePath
}

func GetConfigPath() string {
	return Config.configPath
}

func GetConfigFile() string {
	return Config.configFile
}
