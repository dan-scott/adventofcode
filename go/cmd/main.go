package main

import (
	"log"

	"gitlab.com/danscott/adventofcode/go/ui"
)

func main() {
	p := ui.NewProgram()

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
