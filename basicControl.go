package gonsole

import "github.com/nsf/termbox-go"

// Control is the base model for a UI control
type BasicControl struct {
	window     *Window
	parent     Control
	id         string
	Position   Box
	Visible    bool
	Enabled    bool
	focussable bool
	ZIndex     int
	TabIndex   int
	dirty      bool

	// styling
	Foreground termbox.Attribute
	Background termbox.Attribute
	border     LineType
	HAlign     HorizontalAlignment
	VAlign     VerticalAlignment
	Margin     Sides
	Padding    Sides
}

func (ctrl *BasicControl) Init(id string) {
	ctrl.SetID(id)
	ctrl.Pollute()
}

func (ctrl *BasicControl) GetAbsolutePosition() Box {
	if parent := ctrl.Parent(); parent != nil {
		return ctrl.Position.Absolute(parent.GetAbsolutePosition())
	}
	return ctrl.Position
}

func (ctrl *BasicControl) BorderBox() Box {
	return ctrl.GetAbsolutePosition().Minus(ctrl.Margin)
}

func (ctrl *BasicControl) ContentBox() Box {
	// substract padding and margin
	contentBox := ctrl.GetAbsolutePosition().Minus(ctrl.Margin).Minus(ctrl.Padding)
	// substract border if applicable
	if ctrl.HasBorder() {
		contentBox = contentBox.Minus(Sides{1, 1, 1, 1})
	}
	return contentBox
}

func (ctrl *BasicControl) DrawBorder() {
	if ctrl.Border() == LineNone {
		return
	}
	borderForeground := ctrl.Foreground
	if ctrl.Focussed() {
		borderForeground = termbox.ColorYellow
	}
	DrawBorder(ctrl.BorderBox(), ctrl.Border(), borderForeground, ctrl.Background)
}

func (ctrl *BasicControl) ParseEvent(ev *termbox.Event) bool {
	// to be implemented for individual controls
	return false
}

func (ctrl *BasicControl) ID() string {
	return ctrl.id
}

func (ctrl *BasicControl) SetID(id string) {
	ctrl.id = id
}

func (ctrl *BasicControl) Dirty() bool {
	return ctrl.dirty
}

func (ctrl *BasicControl) Pollute() {
	ctrl.dirty = true
}

func (ctrl *BasicControl) SetWindow(win *Window) {
	ctrl.window = win
}

func (ctrl *BasicControl) Border() LineType {
	return ctrl.border
}

func (ctrl *BasicControl) SetBorder(border LineType) {
	ctrl.border = border
}

func (ctrl *BasicControl) HasBorder() bool {
	return ctrl.border != LineNone
}

func (ctrl *BasicControl) AddEventListener(eventType string, handler func(ev *Event) bool) {
	ctrl.Window().App.EventDispatcher.AddEventListener(ctrl, eventType, handler)
}

func (ctrl *BasicControl) SubmitEvent(ev *Event) {
	ctrl.Window().App.EventDispatcher.SubmitEvent(ev)
}

func (ctrl *BasicControl) Repaint() {
	if !ctrl.Dirty() {
		return
	}
	ClearRect(ctrl.BorderBox(), termbox.ColorDefault, termbox.ColorDefault)
	if ctrl.Background != 0 {
		FillRect(ctrl.Position, ctrl.Foreground, ctrl.Background)
	}
	ctrl.DrawBorder()
	// implement details in controls
}

func (ctrl *BasicControl) Focussed() bool {
	return ctrl.Window().FocussedControl().ID() == ctrl.ID()
}

func (ctrl *BasicControl) Focus() {
	ctrl.Window().SetFocussedControl(ctrl)
}

func (ctrl *BasicControl) Focussable() bool {
	return ctrl.focussable
}

func (ctrl *BasicControl) SetFocussable(focussable bool) {
	ctrl.focussable = focussable
}

func (ctrl *BasicControl) Parent() Control {
	return ctrl.parent
}

func (ctrl *BasicControl) SetParent(parent Control) {
	ctrl.parent = parent
}

func (ctrl *BasicControl) Window() *Window {
	if win := ctrl.window; win != nil {
		return win
	}
	if parent := ctrl.Parent(); parent != nil {
		return parent.Window()
	}
	return nil
}
