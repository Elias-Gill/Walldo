/*
    Paquete que se encarga de leer, administrar y actualizar las configuraciones
del usuario
*/
package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	global "github.com/elias-gill/walldo-in-go/globals"
)

// Return all folders configured by the user in the configuration file
func ConfiguredPaths() []string {
	return readConfigFile()["Paths"]
}

// Reads and parse the config file
func readConfigFile() map[string][]string {
	// Si no se encuentra el archivo de configuracion entonces lo crea
	fileContent, err := os.Open(global.ConfigDir)
	if err != nil {
		fileContent = crearConfig(global.ConfigDir, global.ConfigPath)
	}
	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var res map[string][]string
	json.Unmarshal([]byte(byteResult), &res)

	return res
}

type Path struct {
	Paths []string
}

// Creates the config folder and the config.json if is not created yet
func crearConfig(dir string, path string) *os.File {
	// create the folders
	os.MkdirAll(path, 0o777)
	os.MkdirAll(path+"resized_images", 0o777)
	os.Create(path + "config.json")

	var data *[]byte
	data = setJsonData()
	return writeJsonData(data, dir)
}

// This is for setting the default data of the config.json file
func setJsonData() *[]byte {
	content := Path{Paths: []string{""}}
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
