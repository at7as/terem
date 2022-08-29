package terem

import (
	"strconv"

	"golang.org/x/sys/windows"
)

// Finish ...
func Finish() error {

	Write(EscColorReset)
	return t.out.Flush()

}

// Clear ...
func Clear() {

	Write(EscClearScreen)
	Write(EscCursorMoveHome)
	Write(EscCursorHide)

}

// Cursor ...
func Cursor(x, y int) {

	Write("\u001b[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H")

}

// CursorHide ...
func CursorHide() {

	Write(EscCursorHide)

}

// CursorShow ...
func CursorShow() {

	Write(EscCursorShow)

}

// Write ...
func Write(s string) {

	windows.Write(t.winO, []byte(s))
	// t.out.WriteString(s)

}

// WriteAt ...
func WriteAt(x, y int, s string) {

	Cursor(x, y)
	Write(s)

}

// Style ...
func Style(fg, bg color) {

	if fg+bg != ColorCurrent {
		s := "\u001b["

		if fg != ColorCurrent {
			s += strconv.Itoa(int(fg))
		}

		if bg != ColorCurrent {
			if fg != ColorCurrent {
				s += ";"
			}
			s += strconv.Itoa(int(bg) + 10)
		}

		s += "m"

		Write(s)

	}

}

// Color ...
func Color(c int) color {

	return color(c)

}

// ColorCurrent ...
const ColorCurrent color = 0

// Color ...
const (
	ColorBlack color = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

// ColorBright ...
const (
	ColorBlackBright color = iota + 90
	ColorRedBright
	ColorGreenBright
	ColorYellowBright
	ColorBlueBright
	ColorMagentaBright
	ColorCyanBright
	ColorWhiteBright
)

// ANSI Escape codes
const (
	EscClearScreen    = "\u001b[2J"
	EscCursorMoveHome = "\u001b[H"
	EscCursorHide     = "\u001b[?25l"
	EscCursorShow     = "\u001b[?25h"
	EscColorReset     = "\u001b[0m"
)
