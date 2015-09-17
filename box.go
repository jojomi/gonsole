package gonsole

type Box struct {
	Left   int
	Top    int
	Width  int
	Height int
}

func (b Box) Right() int {
	return b.Left + b.Width - 1
}

func (b Box) Bottom() int {
	return b.Top + b.Height - 1
}
