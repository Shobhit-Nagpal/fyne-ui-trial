package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
  tidyUp()
}

//For a GUI application to work it needs to run an event loop (sometimes called a runloop) 
//that processes user interactions and drawing events.

//In Fyne this is started using the App.Run() or Window.ShowAndRun() functions. 
//One of these must be called from the end of your setup code in the main() function.

func tidyUp() {
	fmt.Println("Exited")
}
