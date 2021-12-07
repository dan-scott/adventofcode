package day07

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"math"
	"sort"
	"strconv"
	"strings"
)

type day07 struct {
	crabs []int64
}

func (d *day07) Open() {
	d.load(inputs.LinesAsString(2021, 7)[0])
}

func (d *day07) Close() {
}

func (d *day07) Part1() string {
	min := int64(math.MaxInt64)
	for pos := d.crabs[0]; pos <= d.crabs[len(d.crabs)-1]; pos++ {
		cost := int64(0)
		for _, crab := range d.crabs {
			cc := pos - crab
			if cc < 0 {
				cc *= -1
			}
			cost += cc
		}
		if cost < min {
			min = cost
		}
	}

	return fmt.Sprint(min)
}

func (d *day07) Part2() string {
	min := int64(math.MaxInt64)
	for pos := d.crabs[0]; pos <= d.crabs[len(d.crabs)-1]; pos++ {
		cost := int64(0)
		for _, crab := range d.crabs {
			cc := pos - crab
			if cc < 0 {
				cc *= -1
			}
			cost += (cc * (cc + 1)) / 2
		}
		if cost < min {
			min = cost
		}
	}

	return fmt.Sprint(min)
}

func (d *day07) load(s string) {
	is := strings.Split(s, ",")
	crabs := make([]int64, len(is))
	for i, v := range is {
		crabs[i], _ = strconv.ParseInt(v, 10, 64)
	}
	sort.SliceStable(crabs, func(i int, j int) bool {
		return crabs[i] < crabs[j]
	})
	d.crabs = crabs
}

func New() runner.Day {
	return &day07{}
}
