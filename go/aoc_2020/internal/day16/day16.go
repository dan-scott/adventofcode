package day16

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"strconv"
	"strings"
)

type day16 struct {
	rules    []rule
	myTicket []int64
	tickets  [][]int64
}

type rule struct {
	name                   string
	min1, max1, min2, max2 int64
}

func (r rule) isValid(val int64) bool {
	return (val >= r.min1 && val <= r.max1) || (val >= r.min2 && val <= r.max2)
}

func parseRule(line string) rule {
	r := rule{}
	p1 := strings.Split(line, ": ")
	r.name = p1[0]
	p2 := strings.Split(p1[1], " or ")
	r.min1, r.max1 = parseRange(p2[0])
	r.min2, r.max2 = parseRange(p2[1])
	return r
}

func parseRange(seg string) (min int64, max int64) {
	p := strings.Split(seg, "-")
	min, _ = strconv.ParseInt(p[0], 10, 64)
	max, _ = strconv.ParseInt(p[1], 10, 64)
	return
}

func parseTicket(line string) []int64 {
	ticket := make([]int64, 0)
	for _, d := range strings.Split(line, ",") {
		v, _ := strconv.ParseInt(d, 10, 64)
		ticket = append(ticket, v)
	}
	return ticket
}

func (d *day16) Open() {
	d.loadLines(inputs.LinesAsString(2020, 16))
}

func (d *day16) loadLines(lines []string) {
	i := 0
	rules := make([]rule, 0)
	for lines[i] != "" {
		rules = append(rules, parseRule(lines[i]))
		i++
	}
	d.rules = rules
	i += 2
	d.myTicket = parseTicket(lines[i])
	i += 3
	tickets := make([][]int64, 0)
	for i < len(lines) {
		tickets = append(tickets, parseTicket(lines[i]))
		i++
	}
	d.tickets = tickets
}

func (d *day16) Close() {
	d.rules = nil
	d.tickets = nil
	d.myTicket = nil
}

func (d *day16) validValue(val int64) bool {
	for _, r := range d.rules {
		if r.isValid(val) {
			return true
		}
	}
	return false
}

func (d *day16) validTicket(i int) bool {
	for _, v := range d.tickets[i] {
		if d.validValue(v) {
			return true
		}
	}
	return false
}

func (d *day16) Part1() string {
	errRate := int64(0)
	for _, t := range d.tickets {
		for _, v := range t {
			if !d.validValue(v) {
				errRate += v
			}
		}
	}
	return fmt.Sprint(errRate)
}

func (d *day16) Part2() string {
	//fieldCheck := make([][]bool, len(d.rules))
	//for i := 0; i < len(fieldCheck); i++ {
	//	fieldCheck[i] = make([]bool, len(d.myTicket))
	//}
	return ""
}

func New() runner.LegacyDay {
	return &day16{}
}
