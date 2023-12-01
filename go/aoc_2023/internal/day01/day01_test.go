package day01

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"testing"
)

func TestDay01_Part1(t *testing.T) {
	input := inputs.StrScanner(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
	expected := "142"
	assert.Equal(t, expected, New().Part1(input))
}

func TestDay01_Part2(t *testing.T) {
	input := inputs.StrScanner(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)
	expected := "281"
	assert.Equal(t, expected, New().Part2(input))
}
