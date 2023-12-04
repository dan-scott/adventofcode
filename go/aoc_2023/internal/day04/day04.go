package day04

import (
	"bufio"
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"strconv"
)

type day04 struct{}

func (d *day04) Year() uint {
	return 2023
}

func (d *day04) Day() uint {
	return 4
}

func (d *day04) Part1(input *bufio.Scanner) string {
	total := 0
	input.Split(bufio.ScanWords)
	var wins map[int]interface{}
	score := 0
	inWins := true
	for input.Scan() {
		word := input.Text()
		switch word {
		case "Card":
			total += score
			wins = make(map[int]interface{})
			input.Scan()
			inWins = true
		case "|":
			inWins = false
			score = 0
		default:
			val, _ := strconv.ParseInt(word, 10, 32)
			if !inWins {
				if _, ok := wins[int(val)]; ok {
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			} else {
				wins[int(val)] = nil
			}
		}
	}
	return fmt.Sprint(total)
}

func (d *day04) Part2(input *bufio.Scanner) string {
	input.Split(bufio.ScanWords)
	cardCounts := make(map[int]int)
	current := 0
	var wins map[int]interface{}
	inWins := true
	winCt := 0
	for input.Scan() {
		word := input.Text()
		switch word {
		case "Card":
			if winCt > 0 {
				for i := current + 1; i <= current+winCt; i++ {
					cardCounts[i] += cardCounts[current]
				}
			}
			input.Scan()
			cardNumStr := input.Text()
			cardNumStr = cardNumStr[:len(cardNumStr)-1]
			cn, _ := strconv.ParseInt(cardNumStr, 10, 32)
			current = int(cn)
			cardCounts[current] += 1
			wins = make(map[int]interface{})
			inWins = true
		case "|":
			inWins = false
			winCt = 0

		default:
			val, _ := strconv.ParseInt(word, 10, 32)
			if !inWins {
				if _, ok := wins[int(val)]; ok {
					winCt += 1
				}
			} else {
				wins[int(val)] = nil
			}
		}
	}
	count := 0
	for i, c := range cardCounts {
		if i > current {
			break
		}
		count += c
	}
	return fmt.Sprint(count)
}

func New() runner.Day {
	return &day04{}
}
