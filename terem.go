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
	color uint8
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
	t.resize()
	t.c = c
	t.out = bufio.NewWriter(os.Stdout)

	return nil

}

// Do ...
func Do(v bool) {

	t.do = v

}

// Run ...
func Run(c Controller) error {

	if t == nil {
		Init(c)
	}

	go Read()

	Do(true)
	for t.do {

		Clear()
		Style(ColorWhite, ColorBlack)

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
