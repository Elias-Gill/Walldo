package explorer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
)

func ConfigWindow(win *fyne.Window) {
    vbox := container.New(layout.NewVBoxLayout())
    dia := dialog.NewCustom("Settings", "close", vbox, *win)
    dia.Show()
}
