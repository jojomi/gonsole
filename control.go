package gonsole

import "github.com/nsf/termbox-go"

type Control interface {
	ID() string
	SetID(id string)

	Focussed() bool
	Focus()

	SetWindow(win *Window)
	Window() *Window

	Parent() Control
	SetParent(parent Control)

	Border() LineType
	SetBorder(lineType LineType)

	Repaint()

	GetAbsolutePosition() Box

	// return true if event was parsed and should not continue bubbling up
	ParseEvent(ev *termbox.Event) bool
	SubmitEvent(ev *Event)
	AddEventListener(eventType string, handler func(ev *Event) bool)
}
