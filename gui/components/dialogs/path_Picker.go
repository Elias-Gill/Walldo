package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// generates and show a new path picker from the file system.
func NewPathPicker(win fyne.Window, callback func(string)) {
	dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			panic("Error picking the file" + err.Error())
		} else if uri == nil {
            callback("")
		} else {
            callback(uri.Path())
		}
	}, win).Show()
}
