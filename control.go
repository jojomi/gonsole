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
	Margin     Sides
	Padding    Sides
}

func (ctrl *BasicControl) GetBorderBox() Box {
	return ctrl.Position.Minus(ctrl.Margin)
}

func (ctrl *BasicControl) GetContentBox() Box {
	// substract padding and margin
	contentBox := ctrl.Position.Minus(ctrl.Margin).Minus(ctrl.Padding)
	// substract border if applicable
	if ctrl.Border != LineNone {
		contentBox = contentBox.Minus(Sides{1, 1, 1, 1})
	}
	return contentBox
}

func (ctrl *BasicControl) DrawBorder() {
	if ctrl.Border == LineNone {
		return
	}
	DrawRect(ctrl.GetBorderBox(), ctrl.Border, ctrl.Foreground, ctrl.Background)
}

type Control interface {
	Repaint()
}
