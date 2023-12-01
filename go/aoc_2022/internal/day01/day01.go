package day01

import (
	"bufio"
	"fmt"
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

func (d *day01) Year() uint { return 2022 }
func (d *day01) Day() uint  { return 1 }

func (d *day01) Part1(input *bufio.Scanner) string {
	elves := d.getElves(input)

	return fmt.Sprint(elves[len(elves)-1])
}

func (d *day01) Part2(input *bufio.Scanner) string {
	elves := d.getElves(input)
	elves = elves[len(elves)-3:]

	return fmt.Sprint(elves[0] + elves[1] + elves[2])
}

func (d *day01) getElves(s *bufio.Scanner) []int64 {
	var elves []int64
	total := int64(0)
	for s.Scan() {
		l := s.Text()
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
