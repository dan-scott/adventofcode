package runner

import (
	"fmt"
	"time"
)

type Day interface {
	Open()
	Close()
	Part1() string
	Part2() string
}

func Run(days []Day) {
	allStart := time.Now()

	for i, d := range days {

		fmt.Printf("Loading Day %d....", i+1)
		start := time.Now()
		d.Open()
		end := time.Now()
		diff := end.Sub(start)
		fmt.Printf("\t  loaded in %15s\n", diff.String())

		fmt.Printf("\tSolving part 1...")
		start = time.Now()
		res := d.Part1()
		end = time.Now()
		diff = end.Sub(start)
		fmt.Printf(" solved in %15s", diff.String())
		fmt.Printf("%20s\n", res)

		fmt.Printf("\tSolving part 2...")
		start = time.Now()
		res = d.Part2()
		end = time.Now()
		diff = end.Sub(start)
		fmt.Printf(" solved in %15s", diff.String())
		fmt.Printf("%20s\n\n", res)

		d.Close()
	}

	allEnd := time.Now()

	fmt.Printf("\nSolutions completed in %s\n", allEnd.Sub(allStart).String())
}
