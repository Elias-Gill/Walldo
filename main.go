package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/gui"
)

func main() {
	uninstallFlag := flag.Bool("uninstall", false, "Uninstall the executable by deleting itself")
	flag.Parse()

	if *uninstallFlag {
		// Get the path to the current executable
		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Error getting executable path:", err)
			return
		}

		// Attempt to delete the executable
		err = os.Remove(exePath)
		if err != nil {
			fmt.Println("Error deleting executable:", err)
			return
		}

		fmt.Println("Executable uninstalled successfully:", exePath)
		return
	}

	config.Init()

	gui.StartGui()
}
