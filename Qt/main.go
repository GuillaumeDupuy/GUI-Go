package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Salutation")
	window.SetGeometry2(300, 300, 500, 500)

	input1 := widgets.NewQLineEdit(nil)
	input1.SetPlaceholderText("Entrez votre prénom")
	input2 := widgets.NewQLineEdit(nil)
	input2.SetPlaceholderText("Entrez votre nom")

	button := widgets.NewQPushButton2("Saluer", nil)
	label := widgets.NewQLabel2("", nil, 0)

	button.ConnectClicked(func(_ bool) {
		label.SetText("Bonjour, " + input1.Text() + " " + input2.Text())
	})

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(input1, 0, 0)
	layout.AddWidget(input2, 0, 0)
	layout.AddWidget(button, 0, 0)
	layout.AddWidget(label, 0, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	window.Show()

	app.Exec()
}
