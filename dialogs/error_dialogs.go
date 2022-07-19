package dialogs

import (
	"fyne.io/fyne/v2/dialog"
	global "github.com/elias-gill/walldo-in-go/globals"
)

// display a dialog error with the current error
func DisplayError(err error) {
	dia := dialog.NewError(err, global.Window)
	dia.Show()
}
