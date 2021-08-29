package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	aspace_bulk_exporter "github.com/dmnyu/aspace-bulk-exporter"
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
	w.SetContent(canvas.NewText(fmt.Sprintf("ASpace Bulk Exporter, %v", aspace_bulk_exporter.AppVersion), color.Black))
	w.ShowAndRun()
}
