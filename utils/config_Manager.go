package utils

import (
	"encoding/json"
	"os"

	global "github.com/elias-gill/walldo-in-go/globals"
)

type Config struct {
	Paths []string `json:"Paths"`
}

func NewConfig() Config {
	return Config{
		Paths: []string{},
	}
}

func (c Config) WithPaths(paths []string) Config {
	c.Paths = paths
	return c
}

// Return all folders configured by the user in the configuration file.
func GetConfiguredPaths() []string {
	return parseConfigFile().Paths
}

func parseConfigFile() Config {
	// create a new config file if it does not exists.
	fileContent, err := os.Open(global.ConfigFile)
	if err != nil {
		MustWriteConfig(Config{})
	}
	defer fileContent.Close()

	var res Config

	json.NewDecoder(fileContent).Decode(&res)

	return res
}

// Creates the config folder and the config.json if is not created yet.
// This function may pannic if cannot modify/create the configuration file.
func MustWriteConfig(config Config) {
	// create the folders
	// os.MkdirAll(global.ConfigPath, 0o777)
	err := os.MkdirAll(global.ConfigPath+"resized_images", 0o777)
	if err != nil {
		panic("Cannot create configuration path: " + err.Error())
	}

	file, err := os.Create(global.ConfigPath + "config.json")
	if err != nil {
		panic("Cannot create the configuration file: " + err.Error())
	}
	defer file.Close()

	json.NewEncoder(file).Encode(config)
}
