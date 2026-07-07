package main

import (
	"flag"

	"github.com/elias-gill/walldo-in-go/cmd"
	"github.com/elias-gill/walldo-in-go/config"
	"github.com/elias-gill/walldo-in-go/gui"
)

func main() {
	uninstallFlag := flag.Bool("uninstall", false, "Uninstall the executable by deleting itself")
	flag.Parse()

	// !IMPORTANT: init configuration before anything else
	config.Init()

	if *uninstallFlag {
		cmd.Uninstall()
		return
	}

	gui.StartGui()
}
