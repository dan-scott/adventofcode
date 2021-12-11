package runner

import (
	"fmt"
	time "github.com/loov/hrtime"
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
		startLoad := time.Now()
		d.Open()
		fmt.Printf("\t  loaded in %15s \n", time.Since(startLoad).String())

		fmt.Printf("\tSolving part 1...")
		startP1 := time.Now()
		res := d.Part1()
		fmt.Printf(" solved in %15s", time.Since(startP1).String())
		fmt.Printf("%20s\n", res)

		fmt.Printf("\tSolving part 2...")
		startP2 := time.Now()
		res = d.Part2()
		fmt.Printf(" solved in %15s", time.Since(startP2).String())
		fmt.Printf("%20s\n\n", res)

		d.Close()
	}

	fmt.Printf("\nSolutions completed in %s\n", time.Since(allStart).String())
}
