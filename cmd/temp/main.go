package main

import (
	"fmt"
	"os"

	"github.com/at7as/terem"
)

func main() {

	go terem.Read()

	for {

		select {
		case e := <-terem.Event:

			if e.Event[10] == 3 {
				os.Exit(0)
			}

			if e.EventType == 1 && e.Event[0] == 1 {

				i := terem.ToCombo(e)
				fmt.Println(i)
				fmt.Println(string(i.Char))

			}

		}

	}

}
