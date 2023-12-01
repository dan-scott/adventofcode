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
	bin   bool
}

func (d *day07) Open() {
	d.load(inputs.LinesAsString(2021, 7)[0])
}

func (d *day07) Close() {
	d.crabs = nil
}

func (d *day07) Part1() string {
	if d.bin {
		return d.part1Binary()
	}
	return d.part1Initial()
}

func (d *day07) Part2() string {
	return d.part2Initial()
}

func (d *day07) part1Initial() string {
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
		} else {
			break
		}
	}

	return fmt.Sprint(min)
}

func (d *day07) part2Initial() string {
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
		} else {
			break
		}
	}

	return fmt.Sprint(min)
}

func (d *day07) part1Binary() string {
	lIdx := 0
	rIdx := len(d.crabs) - 1
	for lIdx != rIdx {
		mIdx := lIdx + (rIdx-lIdx)/2
		lCost := d.costAt(lIdx)
		rCost := d.costAt(rIdx)
		if lCost > rCost {
			lIdx = mIdx
		} else {
			rIdx = mIdx
		}
	}
	return fmt.Sprint(d.costAt(lIdx))
}

func dst(a, b int64) int64 {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

func sumDst(a, b int64) int64 {
	n := dst(a, b)
	return (n * (n + 1)) / 2
}

func (d *day07) costAt(idx int) int64 {
	cost := int64(0)
	pos := d.crabs[idx]
	for _, crab := range d.crabs {
		cost += dst(pos, crab)
	}
	return cost
}

func (d *day07) sumCostAt(idx int) int64 {
	cost := int64(0)
	pos := d.crabs[idx]
	for _, crab := range d.crabs {
		cost += sumDst(pos, crab)
	}
	return cost
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

func New() runner.LegacyDay {
	return &day07{bin: true}
}
