package ui

import (
	"errors"
	"image"
	"image/png"
	"os"
	"strconv"

	"ztm/pixl/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, err error) {
		if uri == nil {
			return
		}

		if err := png.Encode(uri, app.PxCanvas.PixelData); err != nil {
			dialog.ShowError(err, app.PixlWindow)
			return
		}
		app.State.SetFilePath(uri.URI().Path())
	}, app.PixlWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save as...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
			return
		}

		tryClose := func(f *os.File) {
			if err := f.Close(); err != nil {
				dialog.ShowError(err, app.PixlWindow)
			}
		}

		f, err := os.Create(app.State.FilePath)
		defer tryClose(f)
		if err != nil {
			dialog.ShowError(err, app.PixlWindow)
			return
		}

		if err := png.Encode(f, app.PxCanvas.PixelData); err != nil {
			dialog.ShowError(err, app.PixlWindow)
			return
		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}
			if width <= 0 {
				return errors.New("must be greater than 0")
			}
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widgetFormEntry := widget.NewFormItem("Witdh", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItmes := []*widget.FormItem{widgetFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItmes, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if err := widthEntry.Validate(); err != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixlWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if err := heightEntry.Validate(); err != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixlWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}

				app.PxCanvas.NewDrawing(pixelWidth, pixelHeight)

			}
		}, app.PixlWindow)
	})
}

func BuildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if uri == nil {
				return
			}

			img, _, err := image.Decode(uri)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}

			app.PxCanvas.LoadImage(img)
			app.State.SetFilePath(uri.URI().Path())
			imgColors := util.GetImageColors(img)
			i := 0

			for c := range imgColors {
				if i == len(app.Swatches) {
					break
				}
				app.Swatches[i].SetColor(c)
				i++
			}

		}, app.PixlWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu("File",
		BuildNewMenu(app),
		BuildOpenMenu(app),
		BuildSaveMenu(app),
		BuildSaveAsMenu(app),
	)
}

func SetupMenu(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixlWindow.SetMainMenu(mainMenu)
}
