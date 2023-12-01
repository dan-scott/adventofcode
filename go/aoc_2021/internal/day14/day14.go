package day14

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"math"
)

func New() runner.LegacyDay {
	return &day14{}
}

type day14 struct {
	lines []string
}

func (d *day14) Open() {
	d.lines = inputs.LinesAsString(2021, 14)
}

func (d *day14) Close() {
	d.lines = nil
}

func (d *day14) Part1() string {
	return d.runIterations(10)
}

func (d *day14) Part2() string {
	return d.runIterations(40)
}

func (d *day14) runIterations(count int) string {
	tmpl, rules := d.parseLines()
	pairs := make(map[string]int)
	cts := make(map[byte]int)
	cts[tmpl[0]]++
	for i := 0; i < len(tmpl)-1; i++ {
		pairs[tmpl[i:i+2]]++
		cts[tmpl[i+1]]++
	}

	for i := 0; i < count; i++ {
		next := make(map[string]int)
		for s, c := range pairs {
			n := rules[s]
			cts[n] += c
			next[string([]byte{s[0], n})] += c
			next[string([]byte{n, s[1]})] += c
		}
		pairs = next
	}

	min := math.MaxInt
	max := 0
	for _, c := range cts {
		if min > c {
			min = c
		}
		if max < c {
			max = c
		}
	}
	return fmt.Sprint(max - min)
}

func (d *day14) parseLines() (template string, rules map[string]byte) {
	template = d.lines[0]
	rules = make(map[string]byte, len(d.lines)-2)
	for _, l := range d.lines[2:] {
		rules[l[:2]] = l[6]
	}
	return
}
