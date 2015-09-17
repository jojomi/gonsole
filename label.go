package gonsole

type Label struct {
	BasicControl
	Text string
	//Alignment
}

func NewLabel(id string) *Label {
	label := &Label{}
	return label
}

func (l *Label) Repaint() {
	l.DrawBorder()
	//DrawText()
}
