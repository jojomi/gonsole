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
	focussedControl Control
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

func (win *Window) SetFocussedControl(ctrl Control) {
	for _, loopFC := range win.getControlsDeep() {
		if loopFC.ID() == ctrl.ID() {
			win.focussedControl = loopFC
		}
	}
}

func (win *Window) FocussedControl() Control {
	return win.focussedControl
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

func (win *Window) FocusPrev() {
	win.moveFocus(-1)
}

func (win *Window) FocusNext() {
	win.moveFocus(1)
}

func (win *Window) moveFocus(num int) {
	focusControls := win.getFocusControls()
	currentFocusControl := win.FocussedControl()
	// get focus index
	index := -1
	for i, loopFC := range focusControls {
		if loopFC.ID() == currentFocusControl.ID() {
			index = i
		}
	}
	newIndex := (index + num + len(focusControls)) % len(focusControls)
	if index == -1 {
		newIndex = 0
	}
	newFocusControl := focusControls[newIndex]
	win.SetFocussedControl(newFocusControl)

	// update focus, mark dirty
	currentFocusControl.Pollute()
	newFocusControl.Pollute()
	//win.FullRepaint()
}

func (win *Window) getFocusControls() []Control {
	// TODO order by tabIndex and filter non-focussable controls
	focusControls := []Control{}
	for _, control := range win.getControlsDeep() {
		if control.Focussable() {
			focusControls = append(focusControls, control)
		}
	}
	return focusControls
}

func (win *Window) getControlsDeep() []Control {
	//termbox.Close()
	controls := make([]Control, 0)
	for _, control := range win.controls {
		//fmt.Println(control.ID())
		container, ok := control.(Container)
		//fmt.Println(ok)
		if ok {
			children := container.ChildrenDeep()
			//fmt.Println(len(children))
			for _, child := range children {
				controls = append(controls, child)
			}
		} else {
			controls = append(controls, control)
		}
	}
	return controls
}

// return true if event was parsed and should not continue bubbling up
func (win *Window) ParseEvent(ev *termbox.Event) bool {
	// TODO window level event parsing, support tabbing for changing focus

	// dispatch event to currently focussed control
	if win.FocussedControl().ParseEvent(ev) {
		return true
	}

	// focus navigation events
	// catch tab key if the focussed control did not need it
	// catch arrow keys if the focussed control did not need them
	if ev.Type == termbox.EventKey {
		switch ev.Key {
		case termbox.KeyTab:
			win.FocusNext()
		case termbox.KeyArrowDown:
			win.FocusNext()
		case termbox.KeyArrowUp:
			win.FocusPrev()
		}
	}

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
