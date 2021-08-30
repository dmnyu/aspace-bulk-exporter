package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

/*
1. load a config file
2. Select an environment
3. Select location to export success
4. Select location to export failures
5. Select number of workers
6. Toggle validation
7. Toggle tab-reformatting
*/

func main() {
	a := app.New()
	w := a.NewWindow("ASpace Bulk Exporter")
	w.Resize(fyne.Size{320,240 })
	configLoad := canvas.NewText("Place Holder for Config load", color.Black)
	repoSelect := canvas.NewText("Place Holder for Repo Select", color.Black)
	options := canvas.NewText("Place Holder for Options", color.Black)
	terminal := canvas.NewText("Place Holder for log output", color.Black)
	c := container.New(layout.NewVBoxLayout(), configLoad, repoSelect, options, terminal)
	w.SetContent(c)
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}
