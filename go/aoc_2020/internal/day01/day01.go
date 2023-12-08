package day01

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
)

func New() runner.LegacyDay {
	return &day01{}
}

type day01 struct {
	inpts []int
}

func (d *day01) Open() {
	d.inpts = inputs.LinesAsInt(2020, 1)
}

func (d *day01) Close() {
	d.inpts = nil
}

func (d *day01) Part1() string {

	for i, a := range d.inpts[:len(d.inpts)-1] {
		for j, b := range d.inpts[i+1:] {
			if a+b == 2020 {
				return fmt.Sprintf("%d-%d: %d", i, j, a*b)
			}
		}
	}
	return "not found"
}

func (d *day01) Part2() string {
	for i, a := range d.inpts[:len(d.inpts)-2] {
		for j, b := range d.inpts[i+1 : len(d.inpts)-1] {
			for k, c := range d.inpts[j+1:] {
				if a+b+c == 2020 {
					return fmt.Sprintf("%d-%d-%d: %d", i, j, k, a*b*c)
				}
			}
		}
	}
	return "not found"
}
