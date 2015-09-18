package gonsole

import "github.com/nsf/termbox-go"

// Window is the top-level struct in gonsole library.
type Window struct {
	App        *App
	ID         string
	Width      int
	Height     int
	Background termbox.Attribute
	Foreground termbox.Attribute

	// internal state
	FocussedControl Control
	controls        []Control
	dirty           bool
}

// NewWindow creates a new window for later display
func NewWindow(id string) *Window {
	win := &Window{
		ID: id,
	}
	return win
}

// IsDirty returns true if the window is marked as dirty
func (win *Window) IsDirty() bool {
	if win.dirty {
		return true
	}
	// TODO loop child controls
	return true
}

func (win *Window) Pollute() {
	win.dirty = true
}

// return true if event was parsed and should not continue bubbling up
func (win *Window) ParseEvent(ev *termbox.Event) bool {
	// TODO window level event parsing, support tabbing for changing focus

	// dispatch event to currently focussed control
	if win.FocussedControl.ParseEvent(ev) {
		return true
	}

	// focus navigation events
	// catch tab key if the focussed control did not need it
	// catch arrow keys if the focussed control did not need them

	return false
}

// Repaint the window
func (win *Window) Repaint() {
	// TODO implement
	// loop controls, paint the dirty ones
	for _, ctrl := range win.GetControls() {
		ctrl.Repaint()
	}
	termbox.Flush()
}

func (win *Window) FullRepaint() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	width, height := termbox.Size()
	FillRect(Box{Left: 0, Top: 0, Width: width + 1, Height: height + 1}, win.Foreground, win.Background)
	// TODO repaint all child controls
	termbox.Flush()
}

func (win *Window) AddControl(ctrl Control) {
	ctrl.SetWindow(win)
	win.controls = append(win.controls, ctrl)
}

func (win *Window) GetControls() []Control {
	// TODO order by zIndex
	return win.controls
}
