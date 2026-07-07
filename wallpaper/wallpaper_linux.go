package wallpaper

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

// Global environment detectors.
var desktop = os.Getenv("XDG_CURRENT_DESKTOP")
var displayServer = os.Getenv("XDG_SESSION_TYPE")

// ================================
// Wallpaper engine implementation
// ================================

// AvailableModes returns the supported wallpaper sizing styles based on the current environment.
func AvailableModes() []FillStyle {
	if isWaylandSession(displayServer) {
		return []FillStyle{FILL_ZOOM, FILL_CENTER, FILL_TILE, FILL_ORIGINAL, FILL_SCALE}
	}

	if isGNOMECompliant(desktop) || desktop == "X-Cinnamon" || desktop == "Mate" || desktop == "Deepin" {
		return []FillStyle{FILL_ZOOM, FILL_CENTER, FILL_TILE, FILL_SCALE}
	}

	return []FillStyle{FILL_ZOOM, FILL_CENTER, FILL_TILE, FILL_ORIGINAL, FILL_SCALE}
}

func SetWallpaper(file string, mode FillStyle) error {
	if isWaylandSession(displayServer) {
		return setWayland(file, mode)
	}

	if isGNOMECompliant(desktop) {
		return setGnome(file, mode)
	}

	switch desktop {
	case "KDE":
		return setKDE(file, mode)
	case "X-Cinnamon":
		return setCinnamon(file, mode)
	case "Mate":
		return setMate(file, mode)
	case "XFCE":
		return setXFCE(file, mode)
	case "LXDE", "LXQT":
		return setLxde(file, mode)
	case "Deepin":
		return setDeepin(file, mode)
	default:
		return setFehBg(file, mode)
	}
}

// ================================
// Setters and mode listers
// ================================

func setCinnamon(file string, mode FillStyle) error {
	err := exec.Command("dconf", "write", "/org/cinnamon/desktop/background/picture-options",
		strconv.Quote(getGNOMEString(mode))).Run()
	if err != nil {
		return err
	}
	return exec.Command("dconf", "write", "/org/cinnamon/desktop/background/picture-uri",
		strconv.Quote("file://"+file)).Run()
}

func setDeepin(file string, mode FillStyle) error {
	err := exec.Command("dconf", "write", "/com/deepin/wrap/gnome/desktop/background/picture-options",
		strconv.Quote(getGNOMEString(mode))).Run()
	if err != nil {
		return err
	}
	return exec.Command("dconf", "write", "/com/deepin/wrap/gnome/desktop/background/picture-uri",
		strconv.Quote("file://"+file)).Run()
}

func setFehBg(file string, mode FillStyle) error {
	var fehMode string
	switch mode {
	case FILL_CENTER:
		fehMode = "--bg-center"
	case FILL_SCALE:
		fehMode = "--bg-scale"
	case FILL_ORIGINAL:
		fehMode = "--bg-max"
	case FILL_TILE:
		fehMode = "--bg-tile"
	default:
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_ZOOM\n")
		fehMode = "--bg-fill"
	}
	return exec.Command("feh", fehMode, file).Run()
}

func setGnome(file string, mode FillStyle) error {
	err := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-options", strconv.Quote(getGNOMEString(mode))).Run()
	if err != nil {
		return err
	}
	return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+file)).Run()
}

func setKDE(path string, mode FillStyle) error {
	err := setKDEMode(mode)
	if err != nil {
		return err
	}

	return evalKDE(`
		for (const desktop of desktops()) {
			desktop.currentConfigGroup = ["Wallpaper", "org.kde.image", "General"]
			desktop.writeConfig("Image", ` + strconv.Quote("file://"+path) + `)
		}`)
}

func setLxde(file string, mode FillStyle) error {
	var lxdeMode string
	switch mode {
	case FILL_CENTER, FILL_ORIGINAL:
		lxdeMode = "center"
	case FILL_SCALE:
		lxdeMode = "fit"
	case FILL_TILE:
		lxdeMode = "tile"
	default:
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_ZOOM\n")
		lxdeMode = "fill"
	}

	err := exec.Command("pcmanfm", "--wallpaper-mode", lxdeMode).Run()
	if err != nil {
		return err
	}
	return exec.Command("pcmanfm", "-w", file).Run()
}

// setWayland manages wallpaper execution in Wayland environments using swaybg.
// Process releasing allows the wallpaper to persist independently from the application runtime.
func setWayland(file string, mode FillStyle) error {
	exec.Command("killall", "swaybg").Run()

	configPath := os.ExpandEnv("$HOME/.config/swaybg")
	if err := os.WriteFile(configPath, []byte(file+"\n"), 0644); err != nil {
		return err
	}

	var waylandMode string
	switch mode {
	case FILL_CENTER, FILL_ORIGINAL:
		waylandMode = "center"
	case FILL_SCALE:
		waylandMode = "fit"
	case FILL_TILE:
		waylandMode = "tile"
	default:
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_ZOOM\n")
		waylandMode = "fill"
	}

	cmd := exec.Command("swaybg", "-m", waylandMode, "-i", file)
	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Process.Release()
}

func setMate(file string, mode FillStyle) error {
	err := exec.Command("dconf", "write", "/org/mate/desktop/background/picture-options",
		strconv.Quote(getGNOMEString(mode))).Run()
	if err != nil {
		return err
	}
	return exec.Command("dconf", "write", "/org/mate/desktop/background/picture-filename", strconv.Quote(file)).Run()
}

func setXFCE(file string, mode FillStyle) error {
	err := setXFCEMode(mode)
	if err != nil {
		return err
	}

	desktops, err := getXFCEProps("last-image")
	if err != nil {
		return err
	}

	for _, desktop := range desktops {
		err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", desktop, "--set", file).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// ================================
// Helper functions
// ================================

func isWaylandSession(displayServer string) bool {
	return strings.Contains(strings.ToLower(displayServer), "wayland")
}

func isGNOMECompliant(desktop string) bool {
	return strings.Contains(desktop, "GNOME") || desktop == "Unity" || desktop == "Pantheon"
}

func getGNOMEString(mode FillStyle) string {
	switch mode {
	case FILL_CENTER:
		return "centered"
	case FILL_SCALE:
		return "scaled"
	case FILL_TILE:
		return "wallpaper"
	default:
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_ZOOM\n")
		return "zoom"
	}
}

// setKDEMode configures the display style specifically for plasma shells via dbus.
func setKDEMode(mode FillStyle) error {
	var kdeMode string
	switch mode {
	case FILL_CENTER:
		kdeMode = "6"
	case FILL_ORIGINAL:
		kdeMode = "2"
	case FILL_SCALE:
		kdeMode = "0"
	case FILL_TILE:
		kdeMode = "3"
	default:
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_ZOOM\n")
		kdeMode = "1"
	}

	return evalKDE(`
		for (const desktop of desktops()) {
			desktop.currentConfigGroup = ["Wallpaper", "org.kde.image", "General"]
			desktop.writeConfig("FillMode", ` + kdeMode + `)
		}`)
}

func evalKDE(script string) error {
	return exec.Command("qdbus6", "org.kde.plasmashell", "/PlasmaShell", "org.kde.PlasmaShell.evaluateScript", script).Run()
}

// setXFCEMode updates properties dynamically for all detected XFCE monitors.
func setXFCEMode(mode FillStyle) error {
	styles, err := getXFCEProps("image-style")
	if err != nil {
		return err
	}

	for _, style := range styles {
		err = exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", style, "--set", getXFCEString(mode)).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func getXFCEString(mode FillStyle) string {
	switch mode {
	case FILL_CENTER:
		return "1"
	case FILL_ORIGINAL:
		return "4"
	case FILL_SCALE:
		return "3"
	case FILL_TILE:
		return "2"
	default:
		fmt.Fprintf(os.Stderr, "Wallpaper mode not supported, defaulting to FILL_ZOOM\n")
		return "5"
	}
}

func getXFCEProps(key string) ([]string, error) {
	output, err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--list").Output()
	if err != nil {
		return nil, err
	}

	var desktops []string
	for line := range strings.SplitSeq(strings.Trim(string(output), "\n"), "\n") {
		if path.Base(line) == key {
			desktops = append(desktops, line)
		}
	}
	return desktops, nil
}
