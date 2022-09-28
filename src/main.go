package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World") // window title

	text := widget.NewLabel("Hello World!") // window content
	w.SetContent(text)                      // passing window content
	w.ShowAndRun()
}
