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
	"runtime"
)

// Retorna solo las carpetas configuradas por el usuario
func ConfiguredPaths() []string {
	return readConfigFile()["Paths"]
}

// leer la configuracion del usuario
// La carpeta donde se busca es en ~/.config/walldo/config.json (unix)
// ~/AppData/Local/walldo/config.json (windows)
func readConfigFile() map[string][]string {
	sys_os := runtime.GOOS
	configDir, err := os.UserHomeDir() // home del usuario
	configPath := configDir            // path de las configuraciones

	// determinar la direccion del archivo de config
	switch sys_os {
	case "windows":
		configDir += "/AppData/Local/walldo/config.json"
		configPath += "/AppData/Local/walldo/"

	default:
		// sistemas Unix (Mac y Linux)
		configDir += "/.config/walldo/config.json"
		configPath += "/.config/walldo/"
	}

	// Si no se encuentra el archivo de configuracion entonces lo crea
    fileContent, err := os.Open(configDir)
	if err != nil {
		fileContent = crearConfig(configDir, configPath)
	}
	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var res map[string][]string
	json.Unmarshal([]byte(byteResult), &res)

	return res
}

// crea el arhivo de configuracion por defecto dependiendo del OS
// La data es archivo con "Paths": vacio
func crearConfig(dir string, path string) *os.File {
    // crear las carpetas necesarias
	os.MkdirAll(path, 0777)
	os.MkdirAll(path+"resized_images", 0777)
    os.Create(path+"config.json")

	var data *[]byte
	data = setJsonData()
	return writeJsonData(data, dir)
}

type Path struct {
	Paths []string
}

// retornar la data por defecto del arhivo de configuracion
func setJsonData() *[]byte {
	content := Path{Paths: []string{""}}
	data, err := json.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}
	return &data
}

// crea y escribe el archivo de configuraciones
// Retorna el nuevo archivo abierto
func writeJsonData(data *[]byte, fileName string) *os.File {
	// crear el archivo para escritura
	err := os.WriteFile(fileName, *data, 0777) // escribir
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Open(fileName)
	return file
}
