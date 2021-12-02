package day02

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/base_go/inputs"
	"gitlab.com/danscott/adventofcode/base_go/runner"
)

type day02 struct {
	ops []string
}

func (d *day02) Open() {
	d.ops = inputs.LinesAsString(2021, 2)
}

func (d *day02) Close() {
	d.ops = nil
}

func (d *day02) Part1() string {
	dpt := int64(0)
	hrz := int64(0)
	for _, o := range d.ops {
		val := int64(o[len(o)-1] - '0')
		switch o[0] {
		case 'f':
			hrz += val
		case 'u':
			dpt -= val
		case 'd':
			dpt += val
		}
	}
	return fmt.Sprint(dpt * hrz)
}

func (d *day02) Part2() string {
	dpt := int64(0)
	hrz := int64(0)
	aim := int64(0)
	for _, o := range d.ops {
		val := int64(o[len(o)-1] - '0')
		switch o[0] {
		case 'f':
			hrz += val
			dpt += val * aim
		case 'u':
			aim -= val
		case 'd':
			aim += val
		}
	}
	return fmt.Sprint(dpt * hrz)
}

func New() runner.Day {
	return &day02{}

}
