package main

import (
	"gitlab.com/danscott/adventofcode/go/aoc_2022/internal/day01"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

func main() {
	days := []runner.Day{
		day01.New(),
	}

	runner.Run(days)
}
