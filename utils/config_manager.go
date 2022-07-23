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

// Retorna solo las carpetas configuradas por el usuario
func ConfiguredPaths() []string {
	return readConfigFile()["Paths"]
}

// leer la configuracion del usuario
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

// crea el arhivo de configuracion por defecto dependiendo del OS
// La data es archivo con "Paths": vacio
func crearConfig(dir string, path string) *os.File {
	// crear las carpetas necesarias
	os.MkdirAll(path, 0o777)
	os.MkdirAll(path+"resized_images", 0o777)
	os.Create(path + "config.json")

	var data *[]byte
	data = setJsonData()
	return writeJsonData(data, dir)
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
	err := os.WriteFile(fileName, *data, 0o777) // escribir
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Open(fileName)
	return file
}
