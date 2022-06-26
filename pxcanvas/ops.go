package pxcanvas

import (
	"fyne.io/fyne/v2"
)

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff

	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize++
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize--
		}
	default:
		pxCanvas.PxSize = 10
	}
}
