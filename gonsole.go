package gonsole

import "github.com/nsf/termbox-go"

// App holds the global gonsole state.
type App struct {
	CloseKey     termbox.Key
	windows      []*Window
	activeWindow *Window
}

// NewApp creates a new app
func NewApp() *App {
	app := &App{}
	return app
}

func (app *App) Repaint() {
	// get active window, paint it
	if app.activeWindow.IsDirty() {
		app.activeWindow.Repaint()
	}
}

func (app *App) AddWindow(win *Window) {
	app.windows = append(app.windows, win)

	// first window is automatically activated
	if len(app.windows) == 1 {
		app.ActivateWindow(win)
	}
}

func (app *App) ActivateWindow(win *Window) {
	app.activeWindow = win
}

func (app *App) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	app.Repaint()
mainloop:
	for {
		// poll events
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if app.CloseKey == ev.Key {
				break mainloop
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		app.Repaint()
	}
}
