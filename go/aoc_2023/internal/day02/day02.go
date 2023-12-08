package day02

import (
	"bufio"
	"fmt"
	"github.com/danscott/adventofcode/go/common/runner"
	"strconv"
	"strings"
)

type round struct {
	r int
	g int
	b int
}

func (r *round) lte(o *round) bool {
	return r.r <= o.r && r.g <= o.g && r.b <= o.b
}

func (r *round) pow() int {
	return r.r * r.g * r.b
}

var maxCubes round = round{r: 12, g: 13, b: 14}

type day02 struct {
}

func (d *day02) Year() uint {
	return 2023
}

func (d *day02) Day() uint {
	return 2
}

func (d *day02) Part1(input *bufio.Scanner) string {
	total := 0
	for input.Scan() {
		line := input.Text()
		id, rounds := parseLine(line)
		valid := true
		for _, r := range rounds {
			if !r.lte(&maxCubes) {
				valid = false
			}
		}
		if valid {
			total += id
		}
	}
	return fmt.Sprint(total)
}

func (d *day02) Part2(input *bufio.Scanner) string {
	total := 0
	for input.Scan() {
		line := input.Text()
		_, rounds := parseLine(line)
		max := round{}
		for _, r := range rounds {
			max.r = maxInt(max.r, r.r)
			max.g = maxInt(max.g, r.g)
			max.b = maxInt(max.b, r.b)
		}
		total += max.pow()
	}

	return fmt.Sprint(total)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseLine(line string) (int, []round) {
	gameTokens := strings.Split(line, ": ")
	gameNumber, _ := strconv.ParseInt(gameTokens[0][5:], 10, 32)
	var rounds []round
	for _, rStr := range strings.Split(gameTokens[1], "; ") {
		r := round{}
		for _, p := range strings.Split(rStr, ", ") {
			toks := strings.Split(p, " ")
			ct, _ := strconv.ParseInt(toks[0], 10, 32)
			switch toks[1] {
			case "red":
				r.r = int(ct)
			case "green":
				r.g = int(ct)
			case "blue":
				r.b = int(ct)
			}
		}
		rounds = append(rounds, r)
	}
	return int(gameNumber), rounds
}

func New() runner.Day {
	return &day02{}
}
