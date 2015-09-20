package gonsole

import "github.com/nsf/termbox-go"

type Checkbox struct {
	BasicControl

	// custom
	Text    string
	Checked bool

	// internal
	label *Label
}

func NewCheckbox(id string) *Checkbox {
	// auxiliary label
	label := NewLabel("__lbl_" + id)

	checkbox := &Checkbox{
		label: label,
	}
	checkbox.Init(id)
	checkbox.SetFocussable(true)
	return checkbox
}

func (c *Checkbox) Repaint() {
	if !c.Dirty() {
		return
	}
	c.BasicControl.Repaint()

	// Box
	var icon rune
	if c.Checked {
		icon = '☑'
	} else {
		icon = '☐'
	}
	contentBox := c.ContentBox()
	foreground := c.Foreground
	if c.Focussed() && !c.HasBorder() {
		foreground = termbox.ColorYellow
	}
	termbox.SetCell(contentBox.Left, contentBox.Top, icon, foreground, c.Background)
	// Label
	label := c.label
	label.Text = c.Text
	label.Position = c.ContentBox().Minus(Sides{Left: 2})
	// make sure the label is repainted too
	label.Pollute()
	label.Repaint()
}

func (chk *Checkbox) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEnter:
			fallthrough
		case termbox.KeySpace:
			// change state
			chk.Checked = !chk.Checked
			// events
			if chk.Checked {
				chk.SubmitEvent(&Event{"checked", chk, nil})
			} else {
				chk.SubmitEvent(&Event{"unchecked", chk, nil})
			}
			chk.SubmitEvent(&Event{"changed", chk, nil})
			return true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
