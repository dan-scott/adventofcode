package day06

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
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
		newFish := current[0]
		total += newFish
		copy(current, current[1:])
		current[6] += newFish
		current[8] = newFish
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

func New() runner.LegacyDay {
	return &day06{}
}
