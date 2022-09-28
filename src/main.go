package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
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
	w.SetContent(clock)

	// go routine for time update
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
}
