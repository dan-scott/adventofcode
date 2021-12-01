package day02

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/base_go/inputs"
	"gitlab.com/danscott/adventofcode/base_go/runner"
	"strconv"
	"strings"
)

type day02 struct {
	inputs []policy
}

type policy struct {
	f, s int
	char string
	//test regexp.Regexp
	pass string
}

func (p policy) String() string {
	return fmt.Sprintf("%d-%d %s: %s", p.f, p.s, p.char, p.pass)
}

func (d *day02) Open() {
	lines := inputs.LinesAsString(2020, 2)
	ins := make([]policy, len(lines))
	for i, l := range lines {
		tokens := strings.Split(l, " ")
		nums := strings.Split(tokens[0], "-")
		f, _ := strconv.ParseInt(nums[0], 10, 32)
		s, _ := strconv.ParseInt(nums[1], 10, 32)
		char := tokens[1][0:1]
		ins[i] = policy{int(f), int(s), char, tokens[2]}
	}
	d.inputs = ins
}

func (d *day02) Close() {
	d.inputs = nil
}

func (d *day02) Part1() string {
	valid := 0
	for _, p := range d.inputs {
		ct := strings.Count(p.pass, p.char)
		if ct >= p.f && ct <= p.s {
			valid++
		}
	}
	return fmt.Sprint(valid)
}

func (d *day02) Part2() string {
	valid := 0
	for _, p := range d.inputs {
		c := p.char[0]
		if p.pass[p.f-1] == c && p.pass[p.s-1] == c {
			continue
		} else {
			valid++
		}
	}
	return fmt.Sprint(valid)
}

func New() runner.Day {
	return &day02{}
}
