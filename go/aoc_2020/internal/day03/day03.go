package day03

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

type day03 struct {
	grid []string
	w, h int
}

func (d *day03) Open() {
	lines := inputs.LinesAsString(2020, 3)
	d.grid, d.w, d.h = lines, len(lines[0]), len(lines)
}

func (d *day03) Close() {
	d.grid, d.w, d.h = nil, 0, 0
}

func (d *day03) Part1() string {
	r := 0
	c := 0
	tc := 0
	for r < d.h {
		if d.grid[r][c] == '#' {
			tc++
		}
		r++
		c = (c + 3) % d.w
	}
	return fmt.Sprint(tc)
}

type tracker struct {
	col, row, down, right, treeCt int
	d                             *day03
	done                          bool
}

func (t *tracker) advance() {
	if t.done {
		return
	}
	if t.d.grid[t.row][t.col] == '#' {
		t.treeCt++
	}
	t.row += t.down
	t.col = (t.col + t.right) % t.d.w
	if t.row >= t.d.h {
		t.done = true
	}
}

func (d *day03) tracker(down, right int) *tracker {
	return &tracker{down: down, right: right, d: d}
}

func (d *day03) Part2() string {
	r := 0
	trackers := []*tracker{
		d.tracker(1, 1),
		d.tracker(1, 3),
		d.tracker(1, 5),
		d.tracker(1, 7),
		d.tracker(2, 1),
	}

	for r < d.h {
		for _, t := range trackers {
			t.advance()
		}
		r++
	}

	total := 1
	for _, t := range trackers {
		total *= t.treeCt
	}

	return fmt.Sprint(total)
}

func New() runner.Day {
	return &day03{}
}
