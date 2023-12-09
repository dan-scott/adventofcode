package day07

import (
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay07_Part1(t *testing.T) {
	input := inputs.StrScanner(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)
	expected := "6440"
	assert.Equal(t, expected, New().Part1(input))
}

func TestDay07_Part2(t *testing.T) {
	input := inputs.StrScanner(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)
	expected := "5905"
	assert.Equal(t, expected, New().Part2(input))
}
