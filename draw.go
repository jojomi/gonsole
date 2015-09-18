// Higher level helper functions for termbox
// TODO support theming
package gonsole

import "github.com/nsf/termbox-go"

func DrawRect(box Box, lineType LineType, color, backgroundColor termbox.Attribute) {
	right := box.Right()
	bottom := box.Bottom()
	runes := getLineRunes(lineType)
	// draw box.Top and bottom lines
	for y := box.Top; y < box.Top+box.Height; y = y + box.Height - 1 {
		DrawLineHorizontal(box.Left+1, y, box.Width, lineType, color, backgroundColor)
	}
	// draw box.Left and right lines
	for x := box.Left; x < box.Left+box.Width; x = x + box.Width - 1 {
		DrawLineVertical(x, box.Top+1, box.Height, lineType, color, backgroundColor)
	}
	// draw corners
	termbox.SetCell(box.Left, box.Top, runes[2], color, backgroundColor)
	termbox.SetCell(right, box.Top, runes[3], color, backgroundColor)
	termbox.SetCell(box.Left, bottom, runes[4], color, backgroundColor)
	termbox.SetCell(right, bottom, runes[5], color, backgroundColor)
}

func DrawLineHorizontal(left, top, width int, lineType LineType, color, backgroundColor termbox.Attribute) {
	for x := left; x < left+width-1; x++ {
		termbox.SetCell(x, top, getLineRunes(lineType)[0], color, backgroundColor)
	}
}

func DrawLineVertical(left, top, height int, lineType LineType, color, backgroundColor termbox.Attribute) {
	for y := top; y < top+height-1; y++ {
		termbox.SetCell(left, y, getLineRunes(lineType)[1], color, backgroundColor)
	}
}

func DrawCursor() {
}

// TODO support line breaking for multiline strings
// TODO support alignment
func DrawTextBox(text string, box Box, foreground termbox.Attribute, background termbox.Attribute) {
	// get number of lines to draw
	//
}

func DrawTextSimple(text string, box Box, foreground termbox.Attribute, background termbox.Attribute) {
	for index, char := range text {
		termbox.SetCell(box.Left+index, box.Top, char, foreground, background)
	}
}

func getLineRunes(lineType LineType) []rune {
	// https://en.wikipedia.org/wiki/Box-drawing_character
	var runes []rune
	switch lineType {
	case LineTransparent:
		runes = []rune{' ', ' ', ' ', ' ', ' ', ' '}
	case LineSingle:
		runes = []rune{'─', '│', '┌', '┐', '└', '┘'}
	case LineDouble:
		runes = []rune{'═', '║', '╔', '╗', '╚', '╝'}
	case LineDashed:
		runes = []rune{'╌', '╎', '┌', '┐', '└', '┘'}
	case LineDotted:
		runes = []rune{'┄', '┆', '┌', '┐', '└', '┘'}
	}
	return runes
}
