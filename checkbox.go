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
	contentBox := c.GetContentBox()
	termbox.SetCell(contentBox.Left, contentBox.Top, icon, c.Foreground, c.Background)
	// Label
	label := c.label
	label.Text = c.Text
	label.Position = c.GetContentBox().Minus(Sides{Left: 2})
	label.Repaint()
}
