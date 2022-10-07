package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("The time is: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock") // window title

	text := widget.NewLabel("Hello World!") // window content
	w.SetContent(text)                      // passing window content

	clock := widget.NewLabel(" ")
	updateTime(clock)
	w.SetContent(canvas.NewText("Time:", color.Black))
	w.SetContent(clock)
	w.SetContent(canvas.NewVerticalGradient(color.White, color.Black))

	// go routine for time update
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
}
