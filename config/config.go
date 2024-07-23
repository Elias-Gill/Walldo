package config

import (
	"encoding/json"
	"fmt"
	"os"

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
	WallpfillMode modes.FillStyle `json:"FillStyle"`
	GridSize      GridSize        `json:"GridSize"`
	Paths         []string        `json:"Paths"`

	cachePath  string
	configPath string
	configFile string

	fyneSettings fyne.Settings
	window       fyne.Window
}

var conf Configuration = initConfig()

func initConfig() Configuration {
	cache, err := os.UserCacheDir()
	if err != nil {
		panic("Cannot locate cache directory")
	}

	home, err := os.UserConfigDir()
	if err != nil {
		panic("Cannot locate user config directory")
	}

	aux := Configuration{
		configPath: home + "/walldo/",
		cachePath:  cache + "/walldo/",
		configFile: home + "/walldo/" + "config.json",
	}

	err = os.MkdirAll(aux.configPath, 0o770)
	if err != nil {
		panic("Cannot create config directory " + err.Error())
	}

	err = os.MkdirAll(aux.cachePath, 0o770)
	if err != nil {
		panic("Cannot create cache directory " + err.Error())
	}

	/*
		Merge users configuration. NOTE: as the default values for fill mode and
					    grid size are integers, they have a default value of 0
	*/
	// Open the file if it exists, or create it if it does not exist
	file, err := os.OpenFile(aux.configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(fmt.Sprintf("Cannot open or create config file: %v", err))
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&aux)

	return aux
}

func WriteConfig() {
	// Open the file (create if it doesn't exist, truncate if it does)
	file, err := os.Create(conf.configFile)
	if err != nil {
		fmt.Printf("could not create the configuration file: %s", err.Error())
		return
	}
	defer file.Close()

	// Create a JSON encoder and set it to write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print with indentation

	// Encode the struct to JSON and write it to the file
	if err := encoder.Encode(conf); err != nil {
		fmt.Printf("Could not encode JSON data: %s", err.Error())
	}
}

func SetGridSize(s GridSize) {
	conf.GridSize = s
}

func GetGridSize() GridSize {
	return conf.GridSize
}

func GetPaths() []string {
	return conf.Paths
}

func SetPaths(s []string) {
	conf.Paths = s
}

func SetWallpFillMode(m modes.FillStyle) {
	conf.WallpfillMode = m
	wallpaper.SetMode(m)
}

func GetWallpFillMode() modes.FillStyle {
	return conf.WallpfillMode
}

func GetFyneSettings() fyne.Settings {
	return conf.fyneSettings
}

func SetFyneSettings(s fyne.Settings) {
	conf.fyneSettings = s
}

func SetWindow(w fyne.Window) {
	conf.window = w
}

func GetWindow() fyne.Window {
	return conf.window
}
