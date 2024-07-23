//go:build darwin

package wallpaper

import (
	"os/exec"
	"strconv"

	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
)

// SetFromFile uses AppleScript to tell Finder to set the desktop wallpaper to specified file.
func setFromFile(file string, _ modes.FillStyle) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}
