package day01

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

func New() runner.Day {
	return &day01{}
}

type day01 struct {
	ms []int
}

func (d *day01) Open() {
	d.ms = inputs.LinesAsInt(2021, 1)
}

func (d *day01) Close() {
	d.ms = nil
}

func (d *day01) Part1() string {
	ct := 0
	for i, m := range d.ms[1:] {
		if m > d.ms[i] {
			ct++
		}
	}

	return fmt.Sprint(ct)

}

func (d *day01) Part2() string {
	ct := 0
	for i := 3; i < len(d.ms); i++ {
		if d.sumAt(i) > d.sumAt(i-1) {
			ct++
		}
	}
	return fmt.Sprint(ct)
}

func (d *day01) sumAt(i int) int {
	return d.ms[i-2] + d.ms[i-1] + d.ms[i]
}
