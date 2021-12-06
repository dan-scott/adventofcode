package day06

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

type day06 struct {
	fish []int
}

func (d *day06) Open() {
	d.loadInput(inputs.LinesAsString(2021, 6)[0])
}

func (d *day06) Close() {
	d.fish = nil
}

func (d *day06) Part1() string {
	return d.genFish(80)
}

func (d *day06) Part2() string {
	return d.genFish(256)
}

func (d *day06) genFish(days int) string {
	current := make([]int, 9)
	total := len(d.fish)
	for _, f := range d.fish {
		current[f]++
	}
	for i := 0; i < days; i++ {
		total += current[0]
		current[0], current[1], current[2], current[3], current[4], current[5], current[6], current[7], current[8] = current[1], current[2], current[3], current[4], current[5], current[6], current[7]+current[0], current[8], current[0]
	}
	return fmt.Sprint(total)
}

func (d *day06) loadInput(s string) {
	fish := make([]int, 0, 300)
	for _, f := range s {
		if f != ',' {
			fish = append(fish, int(f-'0'))
		}
	}
	d.fish = fish
}

func New() runner.Day {
	return &day06{}
}
