package explorer

import (
	"fyne.io/fyne/v2"
)

func ConfigWindow(app fyne.App) {
    win := app.NewWindow("Configuration")
    win.Resize(fyne.NewSize(300, 300))
    
}
