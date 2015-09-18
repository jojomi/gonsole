package gonsole

type Label struct {
	BasicControl
	Text string
	//Alignment
}

func NewLabel(id string) *Label {
	label := &Label{}
	label.SetID(id)
	return label
}

func (l *Label) Repaint() {
	l.DrawBorder()
	DrawTextSimple(l.Text, l.GetContentBox(), l.Foreground, l.Background)
}
