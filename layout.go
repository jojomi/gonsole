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

func (b Box) Plus(s Sides) Box {
	return Box{
		Left:   b.Left - s.Left,
		Top:    b.Top - s.Top,
		Width:  b.Width + s.Left + s.Right,
		Height: b.Height + s.Top + s.Bottom,
	}
}

func (b Box) Minus(s Sides) Box {
	return Box{
		Left:   b.Left + s.Left,
		Top:    b.Top + s.Top,
		Width:  b.Width - s.Left - s.Right,
		Height: b.Height - s.Top - s.Bottom,
	}
}

type Sides struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

type LineType int

const (
	LineNone = iota
	LineTransparent
	LineSingle
	LineDouble
	LineDashed
	LineDotted
)

type HorizontalAlignment int

const (
	HorizontalAlignmentLeft = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

type VerticalAlignment int

const (
	HorizontalAlignmentTop = iota
	HorizontalAlignmentMiddle
	HorizontalAlignmentBottom
)
