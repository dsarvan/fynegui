package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("The time is: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("FDTD") // window title
	w.SetContent(canvas.NewText("Finite-difference time-domain method", color.Black))

	// menu list
	file := fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() {
			dialog.ShowFileOpen(func(read fyne.URIReadCloser, err error) {
				fmt.Println("User choose:", read.URI().String(), err)
			}, w)
		}),
		fyne.NewMenuItem("Save", func() {
			dialog.ShowFileSave(func(write fyne.URIWriteCloser, err error) {
				fmt.Println("User choose:", write.URI().String(), err)
			}, w)
		}),
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)

	edit := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)

	view := fyne.NewMenu("View",
		fyne.NewMenuItem("Set default layout", func() {
		}),
	)

	help := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Finite-difference time-domain method"),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Saravanan Dayalan"),
				widget.NewLabel("Licence: MIT"),
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
		edit,
		view,
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

	//w.SetContent(canvas.NewVerticalGradient(color.White, color.Black))

	// go routine for time update
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	acc := widget.NewAccordion(
		widget.NewAccordionItem("Set default layout", widget.NewLabel("Hidden")),
		widget.NewAccordionItem("Windows", widget.NewLabel("Object Library")),
		widget.NewAccordionItem("Toolbars", widget.NewLabel("Hidden")),
	)
	acc.Items[1].Open = true
	w.SetContent(acc)

	// infinite progress bar
	infinite := widget.NewProgressBarInfinite()
	w.SetContent(container.NewVBox(infinite))

	// processor selection
	process := widget.NewRadioGroup([]string{"CPU", "MPI", "GPU"}, func(s string) { log.Println("Selected", s) })

	// dimension selection
	dimension := widget.NewSelect([]string{"1-dimensional", "2-dimensional", "3-dimensional"}, func(s string) { log.Println("Selected", s) })

	vsplit := container.NewVSplit(process, dimension)
	w.SetContent(vsplit)

	// container grid layout
	contain := container.New(layout.NewGridLayoutWithColumns(5), acc, text, clock, vsplit, infinite)
	contain.Resize(fyne.NewSize(120, 120))
	w.SetContent(contain) // display the content

	w.ShowAndRun()
}
