package day05

import (
	"bufio"
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"github.com/danscott/adventofcode/go/common/stack"
	"slices"
	"sort"
)

type day05 struct {
}

func (d *day05) Year() uint {
	return 2023
}

func (d *day05) Day() uint {
	return 5
}

type mapping struct {
	srcStart, srcEnd, destStart int
}

func (m *mapping) contains(val int) bool {
	return val >= m.srcStart && val < m.srcEnd
}

func (m *mapping) getDestVal(val int) int {
	return m.destStart + (val - m.srcStart)
}

func (d *day05) Part1(input *bufio.Scanner) string {
	var seeds []int
	var mappings []mapping
	input.Scan()
	seedScanner := inputs.ScanInts(input.Text()[7:])
	for seedScanner.Scan() {
		seeds = append(seeds, seedScanner.Int())
	}
	for input.Scan() {
		line := input.Text()
		if line == "" {
			next := make([]int, 0, len(seeds))
		Seeds:
			for _, s := range seeds {
				for _, m := range mappings {
					if m.contains(s) {
						next = append(next, m.destStart+(s-m.srcStart))
						continue Seeds
					}
				}
				next = append(next, s)
			}
			copy(seeds, next)
			mappings = make([]mapping, 0)
			if !input.Scan() {
				break
			}
			line = input.Text()
		}
		mapScanner := inputs.ScanInts(line)
		mapScanner.Scan()
		destStart := mapScanner.Int()
		mapScanner.Scan()
		srcStart := mapScanner.Int()
		mapScanner.Scan()
		rangeLen := mapScanner.Int()
		mappings = append(mappings, mapping{
			srcStart:  srcStart,
			srcEnd:    srcStart + rangeLen,
			destStart: destStart,
		})
	}
	next := make([]int, 0, len(seeds))
Seeds2:
	for _, s := range seeds {
		for _, m := range mappings {
			if s >= m.srcStart && s < m.srcEnd {
				next = append(next, m.destStart+(s-m.srcStart))
				continue Seeds2
			}
		}
		next = append(next, s)
	}
	sort.Ints(next)
	return fmt.Sprint(next[0])
}

type seedRange struct {
	start, end int
}

func (d *day05) Part2(input *bufio.Scanner) string {
	var seeds []seedRange
	var mappings []mapping
	input.Scan()
	seedScanner := inputs.ScanInts(input.Text()[7:])
	for {
		if s, ok := seedScanner.NextInt(); ok {
			r, _ := seedScanner.NextInt()
			seeds = append(seeds, seedRange{s, s + r})
		} else {
			break
		}
	}

	input.Scan()
	inMapping := false
	for {
		scanned := input.Scan()
		line := input.Text()
		if !inMapping && line != "" {
			inMapping = true
		} else if inMapping && (line == "" || !scanned) {
			inMapping = false
			rangeStack := stack.New[seedRange](seeds...)
			seeds = []seedRange{}
			for {
				s, ok := rangeStack.Pop()
				if !ok {
					break
				}

				matched := false
				for _, m := range mappings {
					start := m.contains(s.start)
					end := m.contains(s.end - 1)

					shouldBreak := true
					if start && end {
						matched = true
						seeds = append(seeds, seedRange{
							start: m.getDestVal(s.start),
							end:   m.getDestVal(s.end),
						})
					} else if start && !end {
						matched = true
						seeds = append(seeds, seedRange{
							start: m.getDestVal(s.start),
							end:   m.getDestVal(m.srcEnd),
						})
						rangeStack.Push(seedRange{
							start: m.srcEnd,
							end:   s.end,
						})
					} else if !start && end {
						matched = true
						rangeStack.Push(seedRange{
							start: s.start,
							end:   m.srcStart,
						})
						seeds = append(seeds, seedRange{
							start: m.getDestVal(m.srcStart),
							end:   m.getDestVal(s.end),
						})
					} else if s.start < m.srcStart && s.end >= m.srcEnd {
						matched = true
						rangeStack.Push(seedRange{
							start: s.start,
							end:   m.srcStart,
						})
						rangeStack.Push(seedRange{
							start: m.srcEnd,
							end:   s.end,
						})
						seeds = append(seeds, seedRange{
							start: m.getDestVal(m.srcStart),
							end:   m.getDestVal(m.srcEnd),
						})
					} else {
						shouldBreak = false
					}

					if shouldBreak {
						break
					}
				}
				if !matched {
					seeds = append(seeds, s)
				}
			}

			if !scanned {
				break
			}
		} else {
			mappingScanner := inputs.ScanInts(line)
			destStart, _ := mappingScanner.NextInt()
			srcStart, _ := mappingScanner.NextInt()
			rangeLen, _ := mappingScanner.NextInt()
			mappings = append(mappings, mapping{
				srcStart:  srcStart,
				srcEnd:    srcStart + rangeLen,
				destStart: destStart,
			})
		}
	}

	slices.SortFunc(seeds, func(a, b seedRange) int {
		return a.start - b.start
	})

	return fmt.Sprint(seeds[0].start)
}

func New() runner.Day {
	return &day05{}
}
