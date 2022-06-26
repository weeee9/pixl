package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetupMenu(app)
	swatchContainer := BuildSwatches(app)
	colorpicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchContainer, nil, colorpicker, app.PxCanvas)

	app.PixlWindow.SetContent(appLayout)
}
