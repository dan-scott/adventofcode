package day08

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"sort"
	"strings"
)

type entry struct {
	signals []string
	output  []string
}

type day08 struct {
	lines []string
}

func (dy *day08) Open() {
	dy.loadLines(inputs.LinesAsString(2021, 8))
}

func (dy *day08) Close() {
	dy.lines = nil
}

func (dy *day08) Part1() string {
	ct := 0
	for _, l := range dy.lines {
		for _, s := range strings.Split(strings.Split(l, " | ")[1], " ") {
			sl := len(s)
			if sl == 2 || sl == 3 || sl == 4 || sl == 7 {
				ct++
			}
		}
	}
	return fmt.Sprint(ct)
}

func (dy *day08) Part2() string {
	sum := 0
	for _, l := range dy.lines {
		ent := parseEntry(l)
		one := ent.signals[0]
		seven := ent.signals[1]
		four := ent.signals[2]
		eight := ent.signals[9]

		cf := one
		bd := subtract(four, one)
		eg := subtract(subtract(eight, seven), four)

		zero, six, nine := findLenSixes(ent, cf, eg)
		var b string
		if strings.Contains(zero, bd[:1]) {
			b = bd[:1]
		} else {
			b = bd[1:]
		}
		var e string
		if strings.Contains(nine, eg[:1]) {
			e = eg[1:]
		} else {
			e = eg[:1]
		}
		two, three, five := findLenFives(ent, b, e)

		numMap := make(map[string]int)
		numMap[zero] = 0
		numMap[one] = 1
		numMap[two] = 2
		numMap[three] = 3
		numMap[four] = 4
		numMap[five] = 5
		numMap[six] = 6
		numMap[seven] = 7
		numMap[eight] = 8
		numMap[nine] = 9

		sum += numMap[ent.output[0]]*1000 + numMap[ent.output[1]]*100 + numMap[ent.output[2]]*10 + numMap[ent.output[3]]
	}
	return fmt.Sprint(sum)
}

func findLenSixes(e entry, cf, eg string) (zero, six, nine string) {
	for _, s := range e.signals[6:9] {
		hasEG := contains(s, eg)
		hasCF := contains(s, cf)
		if hasEG && hasCF {
			zero = s
		} else if hasEG && !hasCF {
			six = s
		} else {
			nine = s
		}
	}
	return
}

func findLenFives(ent entry, b, e string) (two, three, five string) {
	for _, s := range ent.signals[3:6] {
		if len(s) != 5 {
			continue
		}
		hasB := strings.Contains(s, b)
		hasE := strings.Contains(s, e)
		if !hasB && !hasE {
			three = s
		} else if hasB && !hasE {
			five = s
		} else {
			two = s
		}
	}
	return
}

func subtract(a, b string) string {
	tgt := a
	for i := 0; i < len(b); i++ {
		tgt = strings.Replace(tgt, b[i:i+1], "", 1)
	}
	return tgt
}

func contains(a, b string) bool {
	return len(a)-len(subtract(a, b)) == len(b)
}

func parseEntry(l string) entry {
	a := strings.Split(l, " | ")
	signals := splitAndSort(a[0])
	sort.SliceStable(signals, func(i int, j int) bool {
		return len(signals[i]) < len(signals[j])
	})
	return entry{
		signals: signals,
		output:  splitAndSort(a[1]),
	}
}

func splitAndSort(part string) []string {
	ls := strings.Split(part, " ")
	for i, s := range ls {
		r := []rune(s)
		sort.SliceStable(r, func(i int, j int) bool { return r[i] < r[j] })
		ls[i] = string(r)
	}
	return ls
}

func (dy *day08) loadLines(lines []string) {
	dy.lines = lines
}

func New() runner.Day {
	return &day08{}
}
