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

func (b Box) Absolute(bParent Box) Box {
	return Box{
		Top:    bParent.Top + b.Top,
		Left:   bParent.Left + b.Left,
		Width:  b.Width,
		Height: b.Height,
	}
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

func (s Sides) Plus(s2 Sides) Sides {
	return Sides{
		Left:   s.Left + s2.Left,
		Top:    s.Top + s2.Top,
		Right:  s.Right + s2.Right,
		Bottom: s.Bottom + s2.Bottom,
	}
}

func (s Sides) Minus(s2 Sides) Sides {
	return Sides{
		Left:   s.Left - s2.Left,
		Top:    s.Top - s2.Top,
		Right:  s.Right - s2.Right,
		Bottom: s.Bottom - s2.Bottom,
	}
}

type LineType int

const (
	LineNone = iota
	LineTransparent
	LineSingle
	LineSingleCorners
	LineDouble
	LineDoubleCorners
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
