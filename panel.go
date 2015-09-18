package gonsole

type Panel struct {
	ContainerControl

	// custom
	Title string
}

func NewPanel(id string) *Panel {
	panel := &Panel{}
	panel.SetID(id)
	return panel
}

func (c *Panel) Repaint() {
	c.ContainerControl.Repaint()
	// content area
	// TODO implement in parent class
}
