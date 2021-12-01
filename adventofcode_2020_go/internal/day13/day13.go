package day13

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/base_go/inputs"
	"gitlab.com/danscott/adventofcode/base_go/runner"
	"strconv"
	"strings"
)

type day13 struct {
	lt    int
	buses []int
}

func (d *day13) Open() {
	d.loadLines(inputs.LinesAsString(2020, 13))
}

func (d *day13) Close() {
	d.buses = nil
}

func (d *day13) loadLines(lines []string) {
	lt, _ := strconv.ParseInt(lines[0], 10, 64)
	d.lt = int(lt)
	ids := make([]int, 0)
	for _, v := range strings.Split(lines[1], ",") {
		id := int64(1)
		if v != "x" {
			id, _ = strconv.ParseInt(v, 10, 64)
		}
		ids = append(ids, int(id))
	}
	d.buses = ids
}

func (d *day13) Part1() string {
	current := d.lt
	for true {
		for _, id := range d.buses {
			if id != 1 && current%id == 0 {
				return fmt.Sprint((current - d.lt) * id)
			}
		}
		current++
	}
	return ""
}

func (d *day13) Part2() string {
	prev, multiplier := d.buses[0], d.buses[0]
	for i, id := range d.buses[1:] {
	Inner:
		for found := prev; ; found += multiplier {
			if (found+i+1)%id == 0 {
				prev = found
				multiplier *= id
				break Inner
			}

		}
	}
	return fmt.Sprint(prev)
}

func New() runner.Day {
	return &day13{}
}
