package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/elias-gill/walldo-in-go/globals"
)

// generates and show a new path picker from the file system.
func NewPathPicker(win fyne.Window, callback func(string)) {
	picker := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			panic("Error picking the file" + err.Error())
		}

		if uri != nil {
			callback(uri.Path())
		}
	}, win)

	picker.Resize(fyne.NewSize(globals.Window.Canvas().Size().Width, globals.Window.Canvas().Size().Height))
	picker.Show()
}
