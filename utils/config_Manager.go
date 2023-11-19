package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	global "github.com/elias-gill/walldo-in-go/globals"
)

// Return all folders configured by the user in the configuration file
func GetConfiguredPaths() []string {
	return readConfigFile()["Paths"]
}

// Reads and parse the config file
func readConfigFile() map[string][]string {
	// Si no se encuentra el archivo de configuracion entonces lo crea
	fileContent, err := os.Open(global.ConfigFile)
	if err != nil {
		fileContent = SetConfig("")
	}
	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)

	var res map[string][]string
	json.Unmarshal([]byte(byteResult), &res)

	return res
}

type Path struct {
	Paths []string
}

// Creates the config folder and the config.json if is not created yet
func SetConfig(paths string) *os.File {
	// create the folders
	// os.MkdirAll(global.ConfigPath, 0o777)
	os.MkdirAll(global.ConfigPath+"resized_images", 0o777)
	os.Create(global.ConfigPath + "config.json")

	var data *[]byte
	if paths != "" {
		data = setJsonData(paths)
	} else {
		data = setJsonData("")
	}
	return writeJsonData(data, global.ConfigFile)
}

// This is for setting the default data of the config.json file
func setJsonData(paths string) *[]byte {
	// split paths
	aux := strings.Split(paths, ",")

	// Trim all white spaces
	for x, e := range aux {
		aux[x] = strings.TrimSpace(e)
	}

	// Delete all backslashes (necesary due to a weird problem with fyne inputs)
	for x, i := range aux {
		aux[x] = strings.ReplaceAll(i, `"\`, "")
	}

	// set the paths and the json content
	content := Path{Paths: aux}
	data, err := json.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}
	return &data
}

// Write the new file with the data specified.
// Returns the config file opened.
func writeJsonData(data *[]byte, fileName string) *os.File {
	// create the file and write
	err := os.WriteFile(fileName, *data, 0o777)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Open(fileName)
	return file
}
