package main

import (
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day01"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day02"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day03"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day04"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day05"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day06"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day07"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day08"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day09"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day10"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day11"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day12"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day13"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day14"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day15"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal/day16"
	"gitlab.com/danscott/adventofcode/go/common/runner"
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
		day08.New(),
		day09.New(),
		day10.New(),
		day11.New(),
		day12.New(),
		day13.New(),
		day14.New(),
		day15.New(),
		day16.New(),
	}

	runner.Run(days)
}
