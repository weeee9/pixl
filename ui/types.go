package ui

import (
	"ztm/pixl/apptype"
	"ztm/pixl/pxcanvas"
	"ztm/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PxCanvas   *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
