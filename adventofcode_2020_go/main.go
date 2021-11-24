package main

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day01"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day02"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day03"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day04"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day05"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day06"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day07"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day08"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day09"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal/day10"
	"time"
)

func main() {
	days := []internal.Day{
		day01.New(),
		day02.New(),
		day03.New(),
		day04.New(),
		day05.New(),
		day06.New(),
		day07.New(),
		day08.New(),
		day09.New(),
		day10.New(),
	}

	allStart := time.Now()

	for i, d := range days {

		fmt.Printf("Loading Day %d....", i+1)
		start := time.Now()
		d.Open()
		end := time.Now()
		diff := end.Sub(start)
		fmt.Printf("\t  loaded in %10dns (%dms)\n", diff.Nanoseconds(), diff.Milliseconds())

		fmt.Printf("\tSolving part 1...")
		start = time.Now()
		res := d.Part1()
		end = time.Now()
		diff = end.Sub(start)
		fmt.Printf(" solved in %10dns (%dms)", diff.Nanoseconds(), diff.Milliseconds())
		fmt.Printf("\t\t%s\n", res)

		fmt.Printf("\tSolving part 2...")
		start = time.Now()
		res = d.Part2()
		end = time.Now()
		diff = end.Sub(start)
		fmt.Printf(" solved in %10dns (%dms)", diff.Nanoseconds(), diff.Milliseconds())
		fmt.Printf("\t\t%s\n\n", res)

		d.Close()
	}

	allEnd := time.Now()

	fmt.Printf("\nSolutions completed in %dms\n", allEnd.Sub(allStart).Milliseconds())

}
