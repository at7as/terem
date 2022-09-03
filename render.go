package terem

import (
	"strconv"
)

// Finish ...
func Finish() error {

	if err := Write(EscColorReset); err != nil {
		return err
	}

	return t.out.Flush()

}

// Clear ...
func Clear() error {

	if err := Write(EscCursorMoveHome); err != nil {
		return err
	}

	if err := Write(EscCursorHide); err != nil {
		return err
	}

	return Write(EscClearScreen)

}

// Move ...
func Move(x, y int) error {

	return Write("\u001b[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H")

}

// Cursor ...
func Cursor(show bool) error {

	if show {
		return Write(EscCursorShow)
	}

	return Write(EscCursorHide)

}

// Write ...
func Write(s string) error {

	if _, err := t.out.WriteString(s); err != nil {
		return err
	}

	return nil

}

// WriteAt ...
func WriteAt(x, y int, s string) error {

	if err := Move(x, y); err != nil {
		return err
	}

	return Write(s)

}

// Style ...
func Style(fg, bg Color) error {

	if fg+bg != ColorEmpty {
		s := "\u001b["

		if fg != ColorEmpty {
			s += strconv.Itoa(int(fg))
		}

		if bg != ColorEmpty {
			if fg != ColorEmpty {
				s += ";"
			}
			s += strconv.Itoa(int(bg) + 10)
		}

		s += "m"

		return Write(s)

	}

	return nil

}

// ColorEmpty ...
const ColorEmpty Color = 0

// Color ...
const (
	ColorBlack Color = iota + 30
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
	ColorBlackBright Color = iota + 90
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
