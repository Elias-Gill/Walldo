//go:build windows

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

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

	fmt.Println("Scheduling executable deletion and exiting...")

	// Spawn a background cmd process to wait for this application to exit before deleting the binary
	cmdArgs := []string{
		"/c",
		"timeout", "/t", "2", "/nobreak",
		"&&", "del", "/f", "/q", exePath,
	}

	cmd := exec.Command("cmd.exe", cmdArgs...)

	// Prevent the child process from inheriting windows or blocking the current process from exiting
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x00000008, // DETACHED_PROCESS
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error scheduling self-deletion:", err)
		return
	}

	// Exit immediately to release the file lock on the executable
	os.Exit(0)
}
