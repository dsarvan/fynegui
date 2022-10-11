package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("The time is: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Finite-difference time-domain method") // window title

	// menu list
	file := fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() {
			dialog.ShowFileOpen(func(read fyne.URIReadCloser, err error) {
				fmt.Println("User choose:", read.URI().String(), err)
			}, w)
		}),
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)

	help := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Saravanan Dayalan"),
			), w)
		}),
	)

	license := fyne.NewMenu("License",
		fyne.NewMenuItem("Info", func() {
			dialog.ShowInformation("License",
				"MIT License", w)
		}),
	)

	menu := fyne.NewMainMenu(
		file,
		help,
		license,
	)
	w.SetMainMenu(menu)
	// menu list-ends

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
