package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var rect *canvas.Rectangle
	a := app.NewWithID("i3space")

	// The window title shouldn't be visible in i3, but hey, lets just add it anyway
	w := a.NewWindow("i3space")

	bQ := widget.NewButton("Quit", func() {
		a.Quit()
	})
	bC := widget.NewButton("Color", func() {
		dialog.ShowColorPicker("Color", "Choose Color", func(newColor color.Color) {
			rect.FillColor = newColor
			rect.StrokeColor = newColor
			rect.Refresh()
		}, w)
	})

	defColor := color.RGBA{46, 46, 46, 1}
	rect = canvas.NewRectangle(defColor)
	rect.StrokeColor = defColor
	rect.StrokeWidth = 0.0
	hb := container.NewHBox(bC, bQ)
	vb := container.NewVBox(hb)

	c := container.NewMax(rect, vb)
	w.SetContent(c)
	w.ShowAndRun()
}
