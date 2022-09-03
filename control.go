package terem

import "encoding/binary"

// Read ...
func Read() {

	var e InputEvent

	for {

		t.err = readConsoleInput(&e)
		Event <- e

	}

}

// Combine ...
func Combine(e InputEvent) InputComb {

	c := InputComb{}
	if e.Event[0] == 1 {
		c.Pressed = true
	}
	c.Key = InputKey(e.Event[6])
	c.Char = rune(binary.BigEndian.Uint16([]byte{e.Event[11], e.Event[10]}))

	m := uint8(e.Event[12])
	if m&0x01 != 0 || m&0x02 != 0 {
		c.Alt = true
	}
	if m&0x04 != 0 || m&0x08 != 0 {
		c.Ctrl = true
	}
	if m&0x10 != 0 {
		c.Shift = true
	}
	if m&0x20 != 0 {
		c.Num = true
	}
	if m&0x40 != 0 {
		c.Scroll = true
	}
	if m&0x80 != 0 {
		c.Caps = true
	}

	return c

}
