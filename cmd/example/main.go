package main

import (
	"os"

	"github.com/at7as/terem"
)

type controller struct {
	count int
	r     rune
	n     int
	s     string
}

func (c *controller) Render() error {

	terem.WriteAt(2, 6, c.s)
	terem.WriteAt(2, 7, string(c.r))

	for i := 0; i < 8; i++ {
		terem.Style(terem.Color(i+30), terem.ColorBlack)
		terem.WriteAt(2+i, 2, "#")
	}
	for i := 0; i < 8; i++ {
		terem.Style(terem.ColorWhite, terem.Color(i+30))
		terem.WriteAt(2+i, 3, "#")
	}
	for i := 0; i < 8; i++ {
		terem.Style(terem.Color(i+90), terem.ColorBlack)
		terem.WriteAt(2+i, 4, "#")
	}
	for i := 0; i < 8; i++ {
		terem.Style(terem.ColorWhite, terem.Color(i+90))
		terem.WriteAt(2+i, 5, "#")
	}

	return nil

}

func (c *controller) Control(e terem.InputEvent) error {

	if e.Event[10] == 3 {
		os.Exit(0)
	}

	return nil

}

func main() {

	terem.Run(&controller{})

}
