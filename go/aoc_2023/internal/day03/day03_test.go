package day03

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"testing"
)

func TestDay03_Part1(t *testing.T) {
	input := inputs.StrScanner(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)
	expected := "4361"

	assert.Equal(t, expected, New().Part1(input))
}

func TestDay03_Part2(t *testing.T) {
	input := inputs.StrScanner(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)
	expected := "467835"

	assert.Equal(t, expected, New().Part2(input))
}
