package day05

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/base_go/inputs"
	"gitlab.com/danscott/adventofcode/base_go/runner"
	"math"
	"sort"
)

type day05 struct {
	seats []seat
}

type seat struct {
	r, c, id int
}

func (d *day05) Open() {
	lines := inputs.LinesAsString(2020, 5)
	d.loadPasses(lines)
}

func (d *day05) loadPasses(lines []string) {
	d.seats = make([]seat, len(lines))
	for i, l := range lines {
		d.seats[i] = parse(l)
	}
}

func parse(pass string) seat {
	backRow := 127.0
	frontRow := 0.0
	for i := 0; i < 7; i++ {
		switch pass[i] {
		case 'F':
			backRow = math.Floor(frontRow + (backRow-frontRow)/2)
			break
		case 'B':
			frontRow = math.Ceil(frontRow + (backRow-frontRow)/2)
			break
		}
	}
	leftCol := 0.0
	rightCol := 7.0
	for i := 7; i < 10; i++ {
		switch pass[i] {
		case 'L':
			rightCol = math.Floor(leftCol + (rightCol-leftCol)/2)
			break
		case 'R':
			leftCol = math.Ceil(leftCol + (rightCol-leftCol)/2)
		}
	}
	return seat{
		r:  int(frontRow),
		c:  int(leftCol),
		id: int(frontRow*8 + leftCol),
	}
}

func (d *day05) Close() {
	d.seats = nil
}

func (d *day05) Part1() string {
	max := 0
	for _, s := range d.seats {
		if s.id > max {
			max = s.id
		}
	}

	return fmt.Sprint(max)
}

func (d *day05) Part2() string {
	sort.Slice(d.seats, func(i int, j int) bool {
		return d.seats[i].id < d.seats[j].id
	})
	for i, s := range d.seats[:len(d.seats)-1] {
		if d.seats[i+1].id == s.id+2 {
			return fmt.Sprint(s.id + 1)
		}
	}
	return "not found"
}

func New() runner.Day {
	return &day05{}
}
