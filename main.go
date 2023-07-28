package main

import (
	global "github.com/elias-gill/walldo-in-go/globals"
    "github.com/elias-gill/walldo-in-go/gui"
)

func main() {
	// set all global variables for the instance
	global.SetupEnvVariables()
    gui.SetupGui()
	// run app
	global.Window.ShowAndRun()
}
