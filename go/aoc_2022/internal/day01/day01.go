package day01

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"slices"
	"strconv"
)

func New() runner.Day {
	return &day01{}
}

type day01 struct {
	input []string
}

func (d *day01) Open() {
	d.input = inputs.LinesAsString(2022, 1)
}

func (d *day01) Close() {
	d.input = nil
}

func (d *day01) Part1() string {
	elves := d.getElves()

	return fmt.Sprint(elves[len(elves)-1])
}

func (d *day01) Part2() string {
	elves := d.getElves()
	elves = elves[len(elves)-3:]

	return fmt.Sprint(elves[0] + elves[1] + elves[2])
}

func (d *day01) getElves() []int64 {
	var elves []int64
	total := int64(0)
	for _, l := range d.input {
		if len(l) == 0 {
			elves = append(elves, total)
			total = 0
		}
		cals, _ := strconv.ParseInt(l, 10, 64)
		total += cals
	}

	slices.Sort(elves)
	return elves
}
