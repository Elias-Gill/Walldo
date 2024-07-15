package globals

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type FillStyle string
type GridDimension string

const (
	// Grid cards sizes.
	SIZE_DEFAULT GridDimension = "Default"
	SIZE_SMALL   GridDimension = "Small"
	SIZE_LARGE   GridDimension = "Large"

	// wallpaper fill strategies.
	FILL_ZOOM     FillStyle = "Zoom Fill"
	FILL_SCALE    FillStyle = "Scale"
	FILL_CENTER   FillStyle = "Center"
	FILL_ORIGINAL FillStyle = "Original"
	FILL_TILE     FillStyle = "Tile"

	// Config constants.
	WindowWidth  = "WindowWidth"
	WindowHeight = "WindowHeight"
	FillStrategy = "FillStrategy"
	GridSize     = "GridSize"
)

var Sizes map[GridDimension]Size = map[GridDimension]Size{
	SIZE_LARGE:   {Width: 145, Height: 125},
	SIZE_DEFAULT: {Width: 115, Height: 105},
	SIZE_SMALL:   {Width: 90, Height: 80},
}

type App struct {
	App     fyne.App
	Window  fyne.Window
	AppSize Size
	Config  Config
}

// NOTE: this is internal app config related, os is not necessary to put it inside Config.
type Size struct {
	Width  float32
	Height float32
}

// return the current grid size of the current configuration.
func (a App) CurrGridSize() Size {
	return Sizes[a.Config.GridSize]
}

func NewApp() *App {
	app := app.NewWithID("walldo")

	aux := &App{
		App:    app,
		Window: app.NewWindow("Walldo in go"),
		Config: initConfig(),
		AppSize: Size{
			// NOTE: keep in mind float types
			Height: float32(app.Preferences().FloatWithFallback(WindowHeight, 600)),
			Width:  float32(app.Preferences().FloatWithFallback(WindowWidth, 1020)),
		},
	}

	// resize to the last size
	aux.Window.Resize(fyne.NewSize(float32(aux.AppSize.Width), float32(aux.AppSize.Height)))
	aux.Window.CenterOnScreen()

	// save size on close
	aux.App.Lifecycle().SetOnStopped(func() {
		aux.App.Preferences().SetFloat(WindowHeight, float64(aux.Window.Canvas().Size().Height))
		aux.App.Preferences().SetFloat(WindowWidth, float64(aux.Window.Canvas().Size().Width))
	})

	aux.App.Lifecycle().SetOnStopped(func() {
		aux.WriteConfig()
	})

	return aux
}
