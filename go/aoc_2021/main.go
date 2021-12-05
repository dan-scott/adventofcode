package main

import (
	"gitlab.com/danscott/adventofcode/go/aoc_2021/internal/day01"
	"gitlab.com/danscott/adventofcode/go/aoc_2021/internal/day02"
	"gitlab.com/danscott/adventofcode/go/aoc_2021/internal/day03"
	"gitlab.com/danscott/adventofcode/go/aoc_2021/internal/day04"
	"gitlab.com/danscott/adventofcode/go/aoc_2021/internal/day05"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

func main() {
	days := []runner.Day{
		day01.New(),
		day02.New(),
		day03.New(),
		day04.New(),
		day05.New(),
	}

	runner.Run(days)
}
