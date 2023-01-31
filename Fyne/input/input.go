package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Salutation")

	input1 := widget.NewEntry()
	input1.PlaceHolder = "Entrez votre prénom"
	input2 := widget.NewEntry()
	input2.PlaceHolder = "Entrez votre nom"

	label := widget.NewLabel("")

	button := widget.NewButton("Saluer", func() {
		label.SetText("Bonjour, " + input1.Text + " " + input2.Text)
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		input1, input2, button, label))

	w.ShowAndRun()
}
