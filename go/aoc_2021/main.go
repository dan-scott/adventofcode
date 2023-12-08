package main

import (
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day01"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day02"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day03"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day04"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day05"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day06"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day07"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day08"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day09"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day10"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day11"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day12"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day13"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day14"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day15"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day16"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day17"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day18"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day19"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day20"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day21"
	"github.com/danscott/adventofcode/go/aoc_2021/internal/day22"
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
		day17.New(),
		day18.New(),
		day19.New(),
		day20.New(),
		day21.New(),
		day22.New(),
	}

	runner.RunLegacy(days)
}
