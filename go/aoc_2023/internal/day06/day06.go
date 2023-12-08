package day06

import (
	"bufio"
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"strconv"
	"strings"
)

type day06 struct {
}

func (d *day06) Year() uint {
	return 2023
}

func (d *day06) Day() uint {
	return 6
}

func (d *day06) Part1(input *bufio.Scanner) string {
	input.Scan()
	timeScanner := inputs.ScanInts(input.Text()[9:])
	input.Scan()
	distScanner := inputs.ScanInts(input.Text()[9:])

	total := 0
	for {
		t, ok := timeScanner.NextInt()
		if !ok {
			break
		}
		d, _ := distScanner.NextInt()
		ct := 0
		for i := 1; i < t; i++ {
			dist := i * (t - i)
			if dist > d {
				ct++
			}
		}
		if total == 0 {
			total = ct
		} else {
			total *= ct
		}
	}

	return fmt.Sprint(total)
}

func (d *day06) Part2(input *bufio.Scanner) string {
	input.Scan()
	time, _ := strconv.ParseInt(strings.ReplaceAll(input.Text()[9:], " ", ""), 10, 64)
	input.Scan()
	dist, _ := strconv.ParseInt(strings.ReplaceAll(input.Text()[9:], " ", ""), 10, 64)
	ct := 0

	for i := int64(1); i < time; i++ {
		d := i * (time - i)
		if d > dist {
			ct++
		}
	}

	return fmt.Sprint(ct)
}

func New() runner.Day {
	return &day06{}
}
