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
	checkbox.SetID(id)
	checkbox.SetFocussable(true)
	return checkbox
}

func (c *Checkbox) Repaint() {
	c.DrawBorder()
	// Box
	var icon rune
	if c.Checked {
		icon = '☑'
	} else {
		icon = '▢'
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
	label.Repaint()
}
