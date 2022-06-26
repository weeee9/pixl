package main

import (
	"image/color"
	"ztm/pixl/apptype"
	"ztm/pixl/pxcanvas"
	"ztm/pixl/swatch"
	"ztm/pixl/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("Pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pxCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxCols:       10,
		PxRows:       10,
		PxSize:       30,
	}

	pxCanvas := pxcanvas.NewPxCanvas(&state, pxCanvasConfig)

	appInit := ui.AppInit{
		PxCanvas:   pxCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
