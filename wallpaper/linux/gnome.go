//go:build linux

package linux

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/elias-gill/walldo-in-go/wallpaper/modes"
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

func IsGNOMECompliant(desktop string) bool {
	return strings.Contains(desktop, "GNOME") || desktop == "Unity" || desktop == "Pantheon"
}

func getGNOMEString(mode modes.FillStyle) string {
	switch mode {
	case modes.FILL_CENTER:
		return "centered"
	case modes.FILL_ZOOM:
		return "zoom"
	case modes.FILL_SCALE:
		return "scaled"
	case modes.FILL_TILE:
		return "wallpaper"
	default:
		panic("invalid wallpaper mode")
	}
}

func SetGnome(file string, mode modes.FillStyle) error {
	exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-options", strconv.Quote(getGNOMEString(mode))).Run()
	return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+file)).Run()

}
