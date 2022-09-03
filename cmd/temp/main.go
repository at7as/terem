package main

import (
	"fmt"
	"os"

	"github.com/at7as/terem"
)

func main() {

	terem.Init(nil)

	go terem.Read()

	for {

		select {
		case e := <-terem.Event:

			if e.EventType == terem.InputTypeKey {

				k := terem.Combine(e)

				if k.Pressed {

					if k.Ctrl && k.Key == terem.KeyC {
						os.Exit(0)
					}

					fmt.Println(string(k.Char))

				}

			}

		}

	}

}
