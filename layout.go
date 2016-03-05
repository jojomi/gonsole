package gonsole

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

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

func (b Box) Position() Position {
	return Position{
		Left:   strconv.Itoa(b.Left),
		Top:    strconv.Itoa(b.Top),
		Width:  strconv.Itoa(b.Width),
		Height: strconv.Itoa(b.Height),
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

type Position struct {
	Left   string
	Top    string
	Width  string
	Height string
}

func calcPosition(m int, p string) int {
	re := regexp.MustCompile(`^(\d+%?)([+-]\d+)?$`)
	values := re.FindStringSubmatch(p)

	if len(values) < 2 {
		panic(fmt.Sprintf("invalid position '%s'", p))
	}

	value := 0

	if values[1][len(values[1])-1:] == "%" {
		percent, err := strconv.Atoi(values[1][:len(values[1])-1])
		if err != nil || percent > 100 || percent < 0 {
			panic(fmt.Sprintf("invalid percent value in position '%s'", p))
		}

		value = int(math.Ceil(float64(m) * (float64(percent) / 100)))
	} else {
		v, err := strconv.Atoi(values[1])
		if err != nil {
			panic(fmt.Sprintf("invalid value in position '%s'", p))
		}

		value = v
	}

	if values[2] != "" {
		offset, err := strconv.Atoi(values[2][1:])
		if err != nil || offset < 0 {
			panic(fmt.Sprintf("invalid offset value in position '%s'", p))
		}
		if values[2][0] == '+' {
			value += offset
		} else {
			value -= offset
		}
	}

	return value
}

func (p Position) Box(w, h int) Box {
	return Box{
		Left:   calcPosition(w, p.Left),
		Top:    calcPosition(h, p.Top),
		Width:  calcPosition(w, p.Width),
		Height: calcPosition(h, p.Height),
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
