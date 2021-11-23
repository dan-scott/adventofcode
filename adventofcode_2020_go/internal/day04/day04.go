package day04

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/inputs"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal"
	"regexp"
	"strings"
)

type day04 struct {
	entries []entry
}

type entry = map[string]string

func (d *day04) Open() {
	entries := make([]entry, 0)

	lines := inputs.LinesAsString(4)
	i := 0
	for i < len(lines) {
		current := make(entry)
		for i < len(lines) && lines[i] != "" {
			for _, e := range strings.Split(lines[i], " ") {
				parts := strings.Split(e, ":")
				if parts[0] != "cid" {
					current[parts[0]] = parts[1]
				}
			}
			i++
		}
		if len(current) > 0 {
			entries = append(entries, current)
		}
		i++
	}
	d.entries = entries
}

func simpleValid(e entry) bool {
	return len(e) == 7
}

var hcl = regexp.MustCompile(`^#[\da-f]{6}$`)
var pid = regexp.MustCompile(`^\d{9}$`)
var ecl = map[string]interface{}{
	"amb": nil,
	"blu": nil,
	"brn": nil,
	"gry": nil,
	"grn": nil,
	"hzl": nil,
	"oth": nil,
}

func complexValid(e entry) bool {
	if !simpleValid(e) {
		return false
	}

	byr := internal.MustParseUint(e["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr := internal.MustParseUint(e["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr := internal.MustParseUint(e["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := e["hgt"]
	hU := hgt[len(hgt)-2:]
	hV := internal.MustParseUint(hgt[:len(hgt)-2])
	if hU == "cm" {
		if hV < 150 || hV > 193 {
			return false
		}
	} else if hU == "in" {
		if hV < 59 || hV > 76 {
			return false
		}
	} else {
		return false
	}

	if !hcl.MatchString(e["hcl"]) {
		return false
	}

	if _, ok := ecl[e["ecl"]]; !ok {
		return false
	}

	if !pid.MatchString(e["pid"]) {
		return false
	}

	return true
}

func (d *day04) Close() {
	d.entries = nil
}

func (d *day04) Part1() string {
	sum := 0
	for _, e := range d.entries {
		if simpleValid(e) {
			sum++
		}
	}
	return fmt.Sprint(sum)
}

func (d *day04) Part2() string {
	sum := 0
	for _, e := range d.entries {
		if complexValid(e) {
			sum++
		}
	}
	return fmt.Sprint(sum)
}

func New() internal.Day {
	return &day04{}
}
