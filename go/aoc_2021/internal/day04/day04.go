package day04

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"github.com/danscott/adventofcode/go/common/vec2"
	"strconv"
	"strings"
)

type day04 struct {
	nums       []int64
	boardInput []string
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
}

func (d *day04) Part1() string {
	boards := d.getInitialBoards()
	for _, c := range d.nums {
		for i, b := range boards {
			b.call(c)
			if b.hasWon {
				boards[i] = boards[len(boards)-1]
				boards = boards[:len(boards)-1]
				return fmt.Sprint(b.unmarked * c)
			}
		}
	}

	return "no winner!?"
}

func (d *day04) Part2() string {
	var lastWon *board
	lastCall := int64(0)
	boards := d.getInitialBoards()
	next := make([]*board, len(boards))
	for _, c := range d.nums {
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
	d.boardInput = lines[1:]
}

func (d *day04) getInitialBoards() []*board {
	boards := make([]*board, 0)
	lines := make([]string, len(d.boardInput))
	copy(lines, d.boardInput)
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
				valMap[v] = vec2.Of(col, row)
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

func New() runner.LegacyDay {
	return &day04{}
}
