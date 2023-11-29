//go:build linux || darwin

package wallpaper

import (
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
)

// SetFromFile uses AppleScript to tell Finder to set the desktop wallpaper to specified file.
func darwinSetFromFile(file string) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}

func darwinGetCacheDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, "Library", "Caches"), nil
}
