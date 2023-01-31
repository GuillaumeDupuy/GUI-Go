package main
 
// import fyne
import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
)
 
func main() {
    // new app
    a := app.New()
    // New title and window
    w := a.NewWindow("Loading images from internet")
    // resize
    w.Resize(fyne.NewSize(400, 400))
    // loading image form url
    r, _ := fyne.LoadResourceFromURLString("https://picsum.photos/200")
    img := canvas.NewImageFromResource(r)
    w.SetContent(img)
    //
    w.ShowAndRun()
}