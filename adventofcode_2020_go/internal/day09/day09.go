package day09

import (
	"errors"
	"fmt"
	"gitlab.com/danscott/adventofcode/base_go/inputs"
	"gitlab.com/danscott/adventofcode/base_go/runner"
	"math"
)

type day09 struct {
	nums []int
}

func (d *day09) Open() {
	lines := inputs.LinesAsInt(2020, 9)
	d.loadLines(lines)
}

func (d *day09) loadLines(lines []int) {
	d.nums = lines
}

func (d *day09) Close() {
	d.nums = nil
}

func (d *day09) Part1() string {
	invalid, _ := d.findInvalid(25)
	return fmt.Sprint(invalid)
}

func (d *day09) Part2() string {
	invalid, _ := d.findInvalid(25)
	for i := 0; i < len(d.nums); i++ {
		total := d.nums[i]
		j := i
		for j < len(d.nums) && total < invalid {
			j++
			total += d.nums[j]
		}
		if total == invalid {
			min := math.MaxInt
			max := 0
			for _, n := range d.nums[i : j+1] {
				if min > n {
					min = n
				}
				if max < n {
					max = n
				}
			}
			return fmt.Sprint(min + max)
		}
	}
	return "not found"
}

func (d *day09) findInvalid(preamble int) (int, error) {
Outer:
	for i := preamble; i < len(d.nums); i++ {
		for j := i - preamble; j < i+preamble-1; j++ {
			for k := j + 1; k < j+preamble; k++ {
				if d.nums[j] != d.nums[k] && d.nums[j]+d.nums[k] == d.nums[i] {
					continue Outer
				}
			}
		}
		return d.nums[i], nil
	}
	return 0, errors.New("no invalid numbers found")
}

func New() runner.Day {
	return &day09{}
}
