package main

import (
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day01"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day02"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day03"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day04"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day05"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day06"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day07"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day08"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day09"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day10"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day11"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day12"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day13"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day14"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day15"
	"github.com/danscott/adventofcode/go/aoc_2020/internal/day16"
	"github.com/danscott/adventofcode/go/common/runner"
)

func main() {
	days := []runner.LegacyDay{
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

	runner.RunLegacy(days)
}
