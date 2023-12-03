package main

import (
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/walldo-in-go/gui"
)

func main() {
	// set all global variables and run
	globals.SetupEnvVariables()
	gui.SetupGui()
    globals.Window.CenterOnScreen()
	globals.Window.ShowAndRun()
}
