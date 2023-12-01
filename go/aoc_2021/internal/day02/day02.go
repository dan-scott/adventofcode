package day02

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

type day02 struct {
	//ops []string
}

func (d *day02) Open() {
}

func (d *day02) Close() {
}

func (d *day02) Part1() string {
	s, close := inputs.Scanner(2021, 2)
	defer close()
	dpt := int64(0)
	hrz := int64(0)
	for s.Scan() {
		o := s.Text()
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
	s, close := inputs.Scanner(2021, 2)
	defer close()
	dpt := int64(0)
	hrz := int64(0)
	aim := int64(0)
	for s.Scan() {
		o := s.Text()
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

func New() runner.LegacyDay {
	return &day02{}
}
