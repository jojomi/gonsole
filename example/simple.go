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

	panel := gonsole.NewPanel("panel1")
	panel.Position = gonsole.Position{"25%", "25%", "50%", "50%"}
	panel.Title = "Test Controls"
	//panel.TitleAlignment =
	//panel.Background = termbox.ColorWhite
	panel.SetBorder(gonsole.LineDashed)
	win.AddControl(panel)
	//win.Background = termbox.ColorBlue

	ctrl := gonsole.NewLabel("lblStatus")
	ctrl.Position = gonsole.Position{"2", "2", "30", "3"}
	ctrl.Text = "Test"
	ctrl.SetBorder(gonsole.LineSingle)
	ctrl.Margin = gonsole.Sides{0, 1, 0, 1}
	win.AddControl(ctrl)

	ctrlChk := gonsole.NewCheckbox("chkActive")
	ctrlChk.Position = gonsole.Position{"2", "2", "30", "3"}
	ctrlChk.Checked = true
	ctrlChk.Text = "Test"
	ctrlChk.SetBorder(gonsole.LineDouble)
	panel.AddControl(ctrlChk)

	ctrlChk2 := gonsole.NewCheckbox("chkActive2")
	ctrlChk2.Position = gonsole.Position{"2", "7", "30", "3"}
	ctrlChk2.Checked = false
	ctrlChk2.Text = "Test with more text"
	panel.AddControl(ctrlChk2)

	ctrlBtn := gonsole.NewButton("MyButton")
	ctrlBtn.Position = gonsole.Position{"2", "10", "40", "3"}
	ctrlBtn.Text = "This is a button. Push me!"
	ctrlBtn.SetBorder(gonsole.LineSingle)
	panel.AddControl(ctrlBtn)

	ctrlBtn2 := gonsole.NewButton("MyButton2")
	ctrlBtn2.Position = gonsole.Position{"2", "14", "40", "3"}
	ctrlBtn2.Text = "This is my second magic button..."
	ctrlBtn2.SetBorder(gonsole.LineSingle)
	panel.AddControl(ctrlBtn2)

	ctrlChk2.Focus()

	app.AddWindow(win)

	// events
	ctrlBtn.AddEventListener("clicked", func(ev *gonsole.Event) bool {
		ctrlBtn.Text = "--- clicked ---"
		return true
	})

	ctrlBtn2.AddEventListener("clicked", func(ev *gonsole.Event) bool {
		btn := ev.Source.(*gonsole.Button)
		btn.Text = "Clicked button"
		return true
	})

	ctrlChk2.AddEventListener("checked", func(ev *gonsole.Event) bool {
		chk := ev.Source.(*gonsole.Checkbox)
		chk.Text = "works"
		return true
	})

	ctrlChk2.AddEventListener("unchecked", func(ev *gonsole.Event) bool {
		chk := ev.Source.(*gonsole.Checkbox)
		chk.Text = "does not work"
		return true
	})

	app.Run()
}
