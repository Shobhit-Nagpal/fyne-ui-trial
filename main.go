package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
  clock := widget.NewLabel("")
  w.SetContent(clock)
  go func() {
    for range time.Tick(time.Second) {
      updateTime(clock)
    }
  }()

	w.ShowAndRun()
  tidyUp()
}

func updateTime(clock *widget.Label) {
  formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

//For a GUI application to work it needs to run an event loop (sometimes called a runloop) 
//that processes user interactions and drawing events.

//In Fyne this is started using the App.Run() or Window.ShowAndRun() functions. 
//One of these must be called from the end of your setup code in the main() function.

func tidyUp() {
	fmt.Println("Exited")
}
