package day08

import (
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay07_Part1(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			expected: "2",
		},
		{
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			expected: "6",
		},
	}

	for _, tc := range cases {
		input := inputs.StrScanner(tc.input)
		assert.Equal(t, tc.expected, New().Part1(input))
	}
}

func TestDay08_Part2(t *testing.T) {
	input := inputs.StrScanner(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`)
	expected := "6"
	assert.Equal(t, expected, New().Part2(input))
}
