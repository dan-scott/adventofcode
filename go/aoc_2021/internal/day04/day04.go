package day04

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"strconv"
	"strings"
)

type day04 struct {
	nums   []int64
	boards []*board
}

type board struct {
	hasWon bool
	values [][]int64
	marked [][]bool
}

func (b *board) reset() {
	for i := range b.marked {
		b.marked[i] = make([]bool, 5)
	}
	b.hasWon = false
}

func (b *board) call(num int64) {
	for y, r := range b.values {
		for x, v := range r {
			if v == num {
				b.marked[y][x] = true
				return
			}
		}
	}
}

func (b *board) won() bool {
	for i := 0; i < 5; i++ {
		hCt := 0
		vCt := 0
		for j := 0; j < 5; j++ {
			if b.marked[i][j] {
				vCt++
			}
			if b.marked[j][i] {
				hCt++
			}
		}
		if vCt == 5 || hCt == 5 {
			b.hasWon = true
			return true
		}
	}
	return false
}

func (b *board) sumUnmarked() int64 {
	sum := int64(0)
	for y, r := range b.marked {
		for x, m := range r {
			if !m {
				sum += b.values[y][x]
			}
		}
	}
	return sum
}

func (d *day04) Open() {
	d.loadLines(inputs.LinesAsString(2021, 4))
}

func (d *day04) Close() {
	d.nums = nil
	d.boards = nil
}

func (d *day04) Part1() string {
	for i, c := range d.nums {
		for _, b := range d.boards {
			b.call(c)
		}
		if i > 3 {
			for _, b := range d.boards {
				if b.won() {
					return fmt.Sprint(b.sumUnmarked() * c)
				}
			}
		}
	}

	return "no winner!?"
}

func (d *day04) Part2() string {
	lastWon := 0
	lastCall := int64(0)
	for _, c := range d.nums {
		for i, b := range d.boards {
			if !b.hasWon {
				b.call(c)
				if b.won() {
					lastWon = i
					lastCall = c
				}
			}
		}
	}
	return fmt.Sprint(d.boards[lastWon].sumUnmarked() * lastCall)
}

func (d *day04) loadLines(lines []string) {
	d.nums = parseNums(lines[0], ",")
	d.boards = parseBoards(lines[1:])
}

func parseBoards(lines []string) []*board {
	boards := make([]*board, 0)
	for len(lines) > 0 {
		bLines := lines[1:6]
		lines = lines[6:]
		values := make([][]int64, 5)
		marked := make([][]bool, 5)
		for i, l := range bLines {
			values[i] = parseNums(l, " ")
			marked[i] = make([]bool, 5)
		}
		boards = append(boards, &board{false, values, marked})
	}
	return boards
}

func parseNums(numLine string, sep string) []int64 {
	ns := strings.Split(strings.ReplaceAll(strings.TrimSpace(numLine), fmt.Sprintf("%s%s", sep, sep), sep), sep)
	nums := make([]int64, len(ns))
	for i, n := range ns {
		nums[i], _ = strconv.ParseInt(n, 10, 64)
	}
	return nums
}

func New() runner.Day {
	return &day04{}
}
