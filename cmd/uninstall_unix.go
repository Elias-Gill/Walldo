//go:build !windows

package cmd

import (
	"fmt"
	"os"

	"github.com/elias-gill/walldo-in-go/config"
)

func Uninstall() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// Remove config folder
	if config.Config.ConfigPath != "" {
		if err := os.RemoveAll(config.Config.ConfigPath); err != nil {
			fmt.Println("Error deleting config directory:", err)
		} else {
			fmt.Println("Config directory removed successfully.")
		}
	}

	// Remove images cache folder
	if config.Config.CachePath != "" {
		if err := os.RemoveAll(config.Config.CachePath); err != nil {
			fmt.Println("Error deleting cache directory:", err)
		} else {
			fmt.Println("Cache directory removed successfully.")
		}
	}

	// On Unix-like systems, the executable can be safely deleted while it is running
	if err := os.Remove(exePath); err != nil {
		fmt.Println("Error deleting executable:", err)
		return
	}

	fmt.Println("Executable uninstalled successfully:", exePath)
	os.Exit(0)
}
