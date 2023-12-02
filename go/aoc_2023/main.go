package main

import (
	"gitlab.com/danscott/adventofcode/go/aoc_2023/internal/day01"
	"gitlab.com/danscott/adventofcode/go/aoc_2023/internal/day02"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

func main() {
	days := []runner.Day{
		day01.New(),
		day02.New(),
	}

	runner.Run(days)
}
