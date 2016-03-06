package gonsole

import "github.com/nsf/termbox-go"

type Control interface {
	ID() string
	SetID(id string)
	Init(id string)

	Focussed() bool
	Focus()
	Focussable() bool
	SetFocussable(focussable bool)

	SetWindow(win *Window)
	Window() *Window

	Parent() Control
	SetParent(parent Control)

	Border() LineType
	SetBorder(lineType LineType)
	HasBorder() bool

	Dirty() bool
	Pollute()
	Repaint()

	GetAbsolutePosition() Box
	ContentBox() Box

	// return true if event was parsed and should not continue bubbling up
	ParseEvent(ev *termbox.Event) bool
	SubmitEvent(ev *Event)
	AddEventListener(eventType string, handler func(ev *Event) bool)
}

type Container interface {
	Control
	Children() []Control
	ChildrenDeep() []Control
}
