package main

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/aoc_2021/internal/day01"
	"gitlab.com/danscott/adventofcode/go/aoc_2021/internal/day02"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"time"
)

func main() {
	days := []runner.Day{
		day01.New(),
		day02.New(),
	}

	runner.Run(days)

	allStart := time.Now()

	for i, d := range days {

		fmt.Printf("Loading Day %d....", i+1)
		start := time.Now()
		d.Open()
		end := time.Now()
		diff := end.Sub(start)
		fmt.Printf("\t  loaded in %10dns (% 6dms)\n", diff.Nanoseconds(), diff.Milliseconds())

		fmt.Printf("\tSolving part 1...")
		start = time.Now()
		res := d.Part1()
		end = time.Now()
		diff = end.Sub(start)
		fmt.Printf(" solved in %10dns (% 6dms)", diff.Nanoseconds(), diff.Milliseconds())
		fmt.Printf("%20s\n", res)

		fmt.Printf("\tSolving part 2...")
		start = time.Now()
		res = d.Part2()
		end = time.Now()
		diff = end.Sub(start)
		fmt.Printf(" solved in %10dns (% 6dms)", diff.Nanoseconds(), diff.Milliseconds())
		fmt.Printf("%20s\n\n", res)

		d.Close()
	}

	allEnd := time.Now()

	fmt.Printf("\nSolutions completed in %dms\n", allEnd.Sub(allStart).Milliseconds())

}
