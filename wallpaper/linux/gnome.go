//go:build linux || darwin
package linux

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/elias-gill/walldo-in-go/globals"
	yaml "gopkg.in/yaml.v2"
)

func removeProtocol(input string) string {
	if len(input) >= 7 && input[:7] == "file://" {
		return input[7:]
	}
	return input
}

func parseDconf(command string, args ...string) (string, error) {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", err
	}

	// unquote string
	var unquoted string
	// the output is quoted with single quotes, which cannot be unquoted using strconv.Unquote, but it is valid yaml
	err = yaml.UnmarshalStrict(output, &unquoted)
	if err != nil {
		return unquoted, err
	}

	return removeProtocol(unquoted), nil
}

func isGNOMECompliant() bool {
	return strings.Contains(Desktop, "GNOME") || Desktop == "Unity" || Desktop == "Pantheon"
}

func (mode Mode) getGNOMEString() string {
	switch mode {
	case globals.FILL_CENTER:
		return "centered"
	case globals.FILL_ZOOM:
		return "zoom"
	case globals.FILL_SCALE:
		return "scaled"
	case globals.FILL_TILE:
		return "wallpaper"
	default:
		panic("invalid wallpaper mode")
	}
}

func setForGnome(file string, mode Mode) error {
	exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-options", strconv.Quote(mode.getGNOMEString())).Run()
	return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+file)).Run()

}
