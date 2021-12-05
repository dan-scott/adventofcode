package day04

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
	"strconv"
	"strings"
)

type day04 struct {
	calls  int
	nums   []int64
	boards []*board
}

type board struct {
	hasWon   bool
	valMap   map[int64]vec2.Vec2
	colCt    []int
	rowCt    []int
	unmarked int64
	marked   [][]bool
}

func (b *board) call(num int64) {
	if pos, ok := b.valMap[num]; ok {
		b.marked[pos.Y][pos.X] = true
		b.unmarked -= num
		b.rowCt[pos.Y]++
		b.colCt[pos.X]++
		if b.rowCt[pos.Y] == 5 {
			b.hasWon = true
		}
		if b.colCt[pos.X] == 5 {
			b.hasWon = true
		}
	}
}

func (d *day04) Open() {
	d.loadLines(inputs.LinesAsString(2021, 4))
}

func (d *day04) Close() {
	d.nums = nil
	d.boards = nil
}

func (d *day04) Part1() string {
	for _, c := range d.nums {
		d.calls++
		for i, b := range d.boards {
			b.call(c)
			if b.hasWon {
				d.boards[i] = d.boards[len(d.boards)-1]
				d.boards = d.boards[:len(d.boards)-1]
				return fmt.Sprint(b.unmarked * c)
			}
		}
	}

	return "no winner!?"
}

func (d *day04) Part2() string {
	var lastWon *board
	lastCall := int64(0)
	boards := make([]*board, len(d.boards))
	copy(boards, d.boards)
	next := make([]*board, len(d.boards))
	for _, c := range d.nums[d.calls:] {
		lostCt := 0
		for _, b := range boards {
			b.call(c)
			if b.hasWon {
				lastWon = b
				lastCall = c
			} else {
				next[lostCt] = b
				lostCt++
			}
		}
		boards = next[:lostCt]
	}
	return fmt.Sprint(lastWon.unmarked * lastCall)
}

func (d *day04) loadLines(lines []string) {
	d.nums = parseNums(lines[0])
	d.boards = parseBoards(lines[1:])
	d.calls = 0
}

func parseBoards(lines []string) []*board {
	boards := make([]*board, 0)
	for len(lines) > 0 {
		bLines := lines[1:6]
		lines = lines[6:]
		valMap := make(map[int64]vec2.Vec2, 25)
		colCt := make([]int, 5)
		rowCt := make([]int, 5)
		marked := make([][]bool, 5)
		unmarked := int64(0)
		for row, l := range bLines {
			for col := 0; col < 5; col++ {
				v, _ := strconv.ParseInt(strings.TrimSpace(l[col*3:col*3+2]), 10, 32)
				unmarked += v
				valMap[v] = vec2.Of(int64(col), int64(row))
			}
			marked[row] = make([]bool, 5)
		}
		boards = append(boards, &board{
			hasWon:   false,
			valMap:   valMap,
			colCt:    colCt,
			rowCt:    rowCt,
			marked:   marked,
			unmarked: unmarked,
		})
	}
	return boards
}

func parseNums(numLine string) []int64 {
	ns := strings.Split(numLine, ",")
	nums := make([]int64, len(ns))
	for i, n := range ns {
		nums[i], _ = strconv.ParseInt(n, 10, 64)
	}
	return nums
}

func New() runner.Day {
	return &day04{}
}
