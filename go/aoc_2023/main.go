package main

import (
	"github.com/danscott/adventofcode/go/aoc_2023/internal/day01"
	"github.com/danscott/adventofcode/go/aoc_2023/internal/day02"
	"github.com/danscott/adventofcode/go/aoc_2023/internal/day03"
	"github.com/danscott/adventofcode/go/aoc_2023/internal/day04"
	"github.com/danscott/adventofcode/go/aoc_2023/internal/day05"
	"github.com/danscott/adventofcode/go/aoc_2023/internal/day06"
	"github.com/danscott/adventofcode/go/aoc_2023/internal/day07"
	"github.com/danscott/adventofcode/go/common/runner"
)

func main() {
	days := []runner.Day{
		day01.New(),
		day02.New(),
		day03.New(),
		day04.New(),
		day05.New(),
		day06.New(),
		day07.New(),
	}

	runner.Run(days)
}
