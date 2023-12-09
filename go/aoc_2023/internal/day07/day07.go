package day07

import (
	"bufio"
	"fmt"
	"github.com/danscott/adventofcode/go/common/runner"
	"slices"
	"strconv"
)

type day07 struct {
}

func (d *day07) Year() uint {
	return 2023
}

func (d *day07) Day() uint {
	return 7
}

func (d *day07) Part1(input *bufio.Scanner) string {
	var hands []*hand
	for input.Scan() {
		hands = append(hands, from1(input.Text()))
	}

	cardVals['J'] = 11
	slices.SortFunc(hands, func(a, b *hand) int {
		return a.cmp(b)
	})

	total := 0
	for rank, hand := range hands {
		total += (rank + 1) * hand.bid
	}

	return fmt.Sprint(total)
}

func (d *day07) Part2(input *bufio.Scanner) string {
	var hands []*hand
	for input.Scan() {
		hands = append(hands, from2(input.Text()))
	}

	cardVals['J'] = 1
	slices.SortFunc(hands, func(a, b *hand) int {
		return a.cmp(b)
	})

	total := 0
	for rank, hand := range hands {
		total += (rank + 1) * hand.bid
	}

	return fmt.Sprint(total)
}

func New() runner.Day {
	return &day07{}
}

type handType = uint

const (
	htHIGH handType = iota
	htONE
	htTWO
	htTHREE
	htFULL
	htFOUR
	htFIVE
)

var cardVals = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type hand struct {
	str string
	bid int
	ht  handType
}

func (h *hand) cmp(b *hand) int {
	if h.ht != b.ht {
		return int(h.ht - b.ht)
	}

	for i := 0; i < 5; i++ {
		diff := cardVals[rune(h.str[i])] - cardVals[rune(b.str[i])]
		if diff != 0 {
			return diff
		}
	}

	panic("You have found identical hands!?")
}

func from1(line string) *hand {
	str := line[0:5]
	bid, _ := strconv.ParseInt(line[6:], 10, 32)
	cards := make(map[rune]int, 0)
	for _, c := range []rune(str) {
		if _, ok := cards[c]; ok {
			cards[c]++
		} else {
			cards[c] = 1
		}
	}
	ht := htHIGH
	switch len(cards) {
	case 1:
		ht = htFIVE
	case 2:
		for _, c := range cards {
			if c == 3 || c == 2 {
				ht = htFULL
			} else {
				ht = htFOUR
			}
			break
		}
	case 3:
		hasPair := false
		for _, c := range cards {
			if c == 2 {
				hasPair = true
			}
		}
		if hasPair {
			ht = htTWO
		} else {
			ht = htTHREE
		}
	case 4:
		ht = htONE
	}

	return &hand{str: str, bid: int(bid), ht: ht}
}

func from2(line string) *hand {
	str := line[0:5]
	bid, _ := strconv.ParseInt(line[6:], 10, 32)
	cards := make(map[rune]int, 0)
	for _, c := range []rune(str) {
		if _, ok := cards[c]; ok {
			cards[c]++
		} else {
			cards[c] = 1
		}
	}
	ht := htHIGH
	switch len(cards) {
	case 1:
		ht = htFIVE
	case 2:
		for _, c := range cards {
			if c == 3 || c == 2 {
				ht = htFULL
			} else {
				ht = htFOUR
			}
			break
		}
	case 3:
		hasPair := false
		for _, c := range cards {
			if c == 2 {
				hasPair = true
			}
		}
		if hasPair {
			ht = htTWO
		} else {
			ht = htTHREE
		}
	case 4:
		ht = htONE
	}

	if jct, ok := cards['J']; ok {
		switch ht {
		case htHIGH:
			ht = htONE
		case htONE:
			ht = htTHREE
		case htTWO:
			if jct == 2 {
				ht = htFOUR
			} else {
				ht = htFULL
			}
		case htTHREE:
			ht = htFOUR
		default:
			ht = htFIVE
		}
	}

	return &hand{str: str, bid: int(bid), ht: ht}
}
