package gonsole

import "github.com/nsf/termbox-go"

// Control is the base model for a UI control
type BasicControl struct {
	ID       string
	Position Box
	Visible  bool
	Enabled  bool
	// styling
	Foreground termbox.Attribute
	Background termbox.Attribute
	Border     LineType
	HAlign     HorizontalAlignment
	VAlign     VerticalAlignment
}

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

func (ctrl *BasicControl) DrawBorder() {
	if ctrl.Border == LineNone {
		return
	}
	DrawRect(ctrl.Position, ctrl.Border, ctrl.Foreground, ctrl.Background)
}

type Control interface {
	Repaint()
}
