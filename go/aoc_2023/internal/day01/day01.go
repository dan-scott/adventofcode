package day01

import (
	"bufio"
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

func New() runner.Day {
	return &day01{}
}

var numberTrie Trie[rune, int]

func init() {
	lk := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	lm := make([]T[rune, int], 0, len(lk))
	for k, v := range lk {
		lm = append(lm, ti([]rune(k), v))
	}
	numberTrie = NewTrie(lm...)
}

type day01 struct{}

func (d *day01) Year() uint {
	return 2023
}

func (d *day01) Day() uint {
	return 1
}

func (d *day01) Part1(input *bufio.Scanner) string {
	total := 0
	for input.Scan() {
		line := input.Text()
		init := -1
		last := 0
		for _, r := range []rune(line) {
			if d, ok := toDigit(r); ok {
				if init < 0 {
					init = d
				}
				last = d
			}
		}
		total += init*10 + last
	}

	return fmt.Sprint(total)
}

func (d *day01) Part2(input *bufio.Scanner) string {
	total := 0
	for input.Scan() {
		line := []rune(input.Text())
		init := -1
		last := 0
		for i := 0; i < len(line); i++ {
			if d, ok := numberTrie.Test(line[i:]); ok {
				if init < 0 {
					init = d
				}
				last = d
			}
		}
		total += init*10 + last
	}
	return fmt.Sprint(total)
}

func toDigit(r rune) (int, bool) {
	if r >= '0' && r <= '9' {
		return int(r - '0'), true
	}
	return 0, false
}
