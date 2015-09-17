package gonsole

import (
	"io/ioutil"
	"strconv"

	"github.com/nsf/termbox-go"
)

// Window is the top-level struct in gonsole library.
type Window struct {
	ID     string
	Width  int
	Height int
	Dirty  bool
	ZIndex int
	//BackgroundColor Color
	//Color Color

	// internal state
	FocussedControl Control
	controls        []Control
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
	if win.Dirty {
		return true
	}
	// TODO loop child controls
	return true
}

// Repaint the window
func (win *Window) Repaint() {
	// TODO implement
	// loop controls, paint the dirty ones
	ioutil.WriteFile("log", []byte(strconv.FormatInt(int64(len(win.GetControls())), 10)), 0640)
	for _, ctrl := range win.GetControls() {
		ctrl.Repaint()
	}
	termbox.Flush()
}

func (win *Window) FullRepaint() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	// TODO repaint all child controls
	termbox.Flush()
}

func (win *Window) AddControl(ctrl Control) {
	win.controls = append(win.controls, ctrl)
}

func (win *Window) GetControls() []Control {
	// TODO order by zIndex
	return win.controls
}
