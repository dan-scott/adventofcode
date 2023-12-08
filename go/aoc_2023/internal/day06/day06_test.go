package day06

import (
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay06_Part1(t *testing.T) {
	input := inputs.StrScanner(`Time:      7  15   30
Distance:  9  40  200`)
	expected := "288"
	assert.Equal(t, expected, New().Part1(input))
}

func TestDay06_Part2(t *testing.T) {
	input := inputs.StrScanner(`Time:      7  15   30
Distance:  9  40  200`)
	expected := "71503"
	assert.Equal(t, expected, New().Part2(input))
}
