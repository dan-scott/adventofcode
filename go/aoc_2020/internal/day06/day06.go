package day06

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
)

type day06 struct {
	groups []group
}

type group struct {
	yeses map[int32]uint8
	size  uint8
}

func (d *day06) Open() {
	lines := inputs.LinesAsString(2020, 6)
	d.loadLines(lines)
}

func (d *day06) Close() {
	d.groups = nil
}

func (d *day06) Part1() string {
	total := 0
	for _, g := range d.groups {
		total += len(g.yeses)
	}
	return fmt.Sprint(total)
}

func (d *day06) Part2() string {
	total := 0
	for _, g := range d.groups {
		for _, c := range g.yeses {
			if c == g.size {
				total++
			}
		}
	}

	return fmt.Sprint(total)
}

func (d *day06) loadLines(lines []string) {
	r := 0
	g := 0
	d.groups = make([]group, 0)
	for r < len(lines) {
		d.groups = append(d.groups, group{yeses: make(map[int32]uint8, 26), size: 0})
		for r < len(lines) && lines[r] != "" {
			for _, c := range lines[r] {
				d.groups[g].yeses[c]++
			}
			d.groups[g].size++
			r++
		}
		r++
		g++
	}
}

func New() runner.LegacyDay {
	return &day06{}
}
