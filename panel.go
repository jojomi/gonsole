package gonsole

type Panel struct {
	ContainerControl

	// custom
	Title string
}

func NewPanel(id string) *Panel {
	panel := &Panel{}
	panel.Init(id)
	return panel
}

func (c *Panel) Repaint() {
	if !c.Dirty() {
		return
	}
	c.ContainerControl.Repaint()

	// draw title
	if c.Title != "" {
		if c.Border() == LineNone {
			c.Padding = c.Padding.Plus(Sides{Top: 1})
		}
		DrawTextSimple(" "+c.Title+" ", c.BorderBox().Minus(Sides{Left: 2}), c.Foreground, c.Background)
	}

	// content area (ContainerControl already takes care of drawing the children)
}
