package terem

import "encoding/binary"

// Read ...
func Read() {

	var e InputEvent

	for {

		if err := readConsoleInput(&e); err != nil {
			panic(err)
		}
		Event <- e

	}

}

// ToCombo ...
func ToCombo(e InputEvent) InputCombo {

	i := InputCombo{}
	i.Key = InputKey(e.Event[6])
	i.Char = rune(binary.BigEndian.Uint16([]byte{e.Event[11], e.Event[10]}))

	m := uint8(e.Event[12])
	if m&0x01 != 0 || m&0x02 != 0 {
		i.Alt = true
	}
	if m&0x04 != 0 || m&0x08 != 0 {
		i.Ctrl = true
	}
	if m&0x10 != 0 {
		i.Shift = true
	}
	if m&0x20 != 0 {
		i.Num = true
	}
	if m&0x40 != 0 {
		i.Scroll = true
	}
	if m&0x80 != 0 {
		i.Caps = true
	}

	return i

}
