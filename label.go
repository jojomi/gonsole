package gonsole

type Label struct {
	BasicControl
	Text string
	//Alignment
}

func NewLabel(id string) *Label {
	label := &Label{}
	label.Init(id)
	return label
}

func (l *Label) Repaint() {
	if !l.Dirty() {
		return
	}
	l.BasicControl.Repaint()

	DrawTextSimple(l.Text, l.ContentBox(), l.Foreground, l.Background)
}
