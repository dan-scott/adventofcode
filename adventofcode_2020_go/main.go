package main

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day01"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day02"
	"time"
)

func main() {
	days := []internal.Day{
		day01.Solver(),
		day02.Solver(),
	}

	for i, d := range days {

		fmt.Printf("Loading Day %d....", i+1)
		start := time.Now()
		d.Open()
		end := time.Now()
		fmt.Printf(" loaded in %dns\n", end.Sub(start).Nanoseconds())

		fmt.Printf("\tSolving part 1...")
		start = time.Now()
		res := d.Part1()
		end = time.Now()
		fmt.Printf(" solved in %dns", end.Sub(start).Nanoseconds())
		fmt.Printf("\t\t%s\n", res)

		fmt.Printf("\tSolving part 2...")
		start = time.Now()
		res = d.Part2()
		end = time.Now()
		fmt.Printf(" solved in %dns", end.Sub(start).Microseconds())
		fmt.Printf("\t\t%s\n\n", res)

		d.Close()
	}
}
