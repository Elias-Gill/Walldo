package globals

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	// internal usage
	ConfigFile string
	ConfigPath string

	// user defined configs
	ThumbnailsPath string
	GridSize       GridDimension
	FillStrategy   FillStyle
	Paths          []string
}

/*
These are the used config files:
- ~/.config/walldo/config.json (unix).
- ~/AppData/Local/walldo/config.json (windows).
*/
func initConfig() Config {
	// set darkmode
	os.Setenv("FYNE_THEME", "dark")

	userConfig, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Cannot establish users home directory: ", err.Error())
	}

	err = os.MkdirAll(userConfig+"walldo/", 0o770)
	if err != nil {
		panic("Cannot create config directory " + err.Error())
	}

	var conf = parseConfigFile(userConfig + "walldo/config.json")

	conf.ConfigPath = userConfig + "walldo/"
	conf.ConfigFile = userConfig + "walldo/config.json"

	err = os.MkdirAll(conf.ThumbnailsPath, 0o770)
	if err != nil {
		panic("Cannot create cache directory " + err.Error())
	}

	return conf
}

// Return all folders configured by the user in the configuration file.
func (a App) GetConfiguredPaths() []string {
	return a.Config.Paths
}

func (a App) WriteConfig() {
	file, err := os.OpenFile(a.Config.ConfigFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(a.Config)
	if err != nil {
		panic(err.Error())
	}
}

// parse configuration file or generate a new one, using the default configuration.
func parseConfigFile(file string) Config {
	fileContent, err := os.Open(file)
	defer fileContent.Close()

	// create a new config file if it does not exists.
	if err != nil {
		file, err := os.Create(file)
		if err != nil {
			panic("Cannot create the configuration file: " + err.Error())
		}
		defer file.Close()
	}

	var conf Config

	json.NewDecoder(fileContent).Decode(&conf)

	// if not custom cache defined use OS default cache directory
	if len(conf.ThumbnailsPath) == 0 {
		cacheDir, err := os.UserCacheDir()
		if err != nil {
			panic("Cannot establish users cache directory: " + err.Error())
		}

		conf.ThumbnailsPath = cacheDir + "/walldo/"
	}

	if len(conf.GridSize) == 0 {
		conf.GridSize = SIZE_DEFAULT
	}

	if len(conf.FillStrategy) == 0 {
		conf.FillStrategy = FILL_ORIGINAL
	}

	return conf
}
