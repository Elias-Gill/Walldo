package wallpaper

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func AvailableModes() []FillStyle {
	return []FillStyle{
		FILL_CENTER,
		FILL_ZOOM,
		FILL_SCALE,
		// FILL_ORIGINAL and FILL_TILE are not supported natively by macOS
	}
}

func SetWallpaper(path string, mode FillStyle) error {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	var displayMode string
	switch mode {
	case FILL_ZOOM:
		displayMode = "fill screen"
	case FILL_SCALE:
		displayMode = "as fit"
	default:
		// Use FILL_CENTER as default to prevent sudden crashes
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_CENTER")
		displayMode = "center"
	}

	// AppleScript that targets both the picture path and its scaling placement
	script := `on run argv
		tell application "System Events"
			set desktopCount to count of desktops
			repeat with i from 1 to desktopCount
				tell desktop i
					set picture to (item 1 of argv)
					set picture display policy to (item 2 of argv)
				end tell
			end repeat
		end tell
	end run`

	return exec.CommandContext(ctx, "osascript", "-e", script, path, displayMode).Run()
}
