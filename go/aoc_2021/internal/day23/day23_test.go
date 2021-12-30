package day23

import (
	"strings"
	"testing"
)

func TestDay23_Part1(t *testing.T) {
	input := strings.Split(`
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`, "\n")[1:]
	d := &day23{input}
	const expected = "12521"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}
