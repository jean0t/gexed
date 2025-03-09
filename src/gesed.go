package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
)

type application struct {
  App fyne.App
  Window fyne.Window
}

func (a *application) setWindow() {
  a.Window = a.App.NewWindow("Gesed")
  a.Window.Resize(fyne.Size{Width: 500, Height: 500})
  a.Window.CenterOnScreen()
}

func (a *application) setContent() {

  //Set the buffer that will be edited
  var Buffer *widget.Entry = widget.NewMultiLineEntry()
  Buffer.Wrapping = fyne.TextWrapWord
  Buffer.OnChanged = func(string) {
    a.Window.SetTitle("Gesed*")
  }

  //TODO: Set the Panel that will show the options


  //Set content is the last thing to be done
  a.Window.SetContent(Buffer)
}

func (a *application) Exec() {
  a.setWindow()
  a.setContent()
  a.Window.ShowAndRun()
}

var App *application = &application{App: app.New()}

func main() {
	App.Exec()
}
