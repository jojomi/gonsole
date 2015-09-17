package main

import (
	"github.com/jojomi/gonsole"
	"github.com/nsf/termbox-go"
)

func main() {
	app := gonsole.NewApp()
	app.CloseKey = termbox.KeyEsc
	//app.CloseKey = 'q'
	win := gonsole.NewWindow("winMain")
	ctrl := gonsole.NewLabel("lblStatus")
	ctrl.Position = gonsole.Box{2, 2, 30, 3}
	ctrl.Text = "Test"
	ctrl.Border = gonsole.LineSingle
	win.AddControl(ctrl)
	app.AddWindow(win)
	app.Run()
}
