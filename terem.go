package terem

import (
	"bufio"
	"os"
)

type (
	terminal struct {
		out  *bufio.Writer
		do   bool
		c    Controller
		size struct {
			w int
			h int
		}
	}
	// Color ...
	Color uint8
	// Controller ...
	Controller interface {
		Render() error
		Control(InputEvent) error
	}
)

func (t *terminal) resize() error {

	w, h, err := getConsoleSize()
	if err != nil {
		return err
	}
	t.size.w = w
	t.size.h = h
	return nil

}

var (
	t *terminal
	// Event ...
	Event = make(chan InputEvent)
)

// Init ...
func Init(c Controller) error {

	t = &terminal{}
	t.c = c
	t.out = bufio.NewWriter(os.Stdout)
	return t.resize()

}

// Do ...
func Do(v bool) {

	t.do = v

}

// Run ...
func Run(c Controller) error {

	if t == nil {
		if err := Init(c); err != nil {
			return err
		}
	}

	go Read(nil)

	Do(true)
	for t.do {

		if err := Clear(); err != nil {
			return err
		}

		if err := Style(ColorWhite, ColorBlack); err != nil {
			return err
		}

		if err := t.c.Render(); err != nil {
			return err
		}

		if err := Finish(); err != nil {
			return err
		}

		select {
		case e := <-Event:
			if err := t.c.Control(e); err != nil {
				return err
			}
		}

	}

	return nil

}

// Size ...
func Size() (w, h int) {

	return t.size.w, t.size.h

}
