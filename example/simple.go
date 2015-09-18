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
	ctrl.Margin = gonsole.Sides{0, 1, 0, 1}
	win.AddControl(ctrl)

	ctrlChk := gonsole.NewCheckbox("chkActive")
	ctrlChk.Position = gonsole.Box{2, 8, 30, 3}
	ctrlChk.Checked = true
	ctrlChk.Text = "Test"
	ctrlChk.Border = gonsole.LineDouble
	win.AddControl(ctrlChk)
	ctrlChk2 := gonsole.NewCheckbox("chkActive")
	ctrlChk2.Position = gonsole.Box{2, 12, 30, 3}
	ctrlChk2.Checked = false
	ctrlChk2.Text = "Test mit mehr Text"
	win.AddControl(ctrlChk2)

	app.AddWindow(win)
	app.Run()
}
