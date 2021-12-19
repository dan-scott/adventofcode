package day17

import "testing"

func TestDay17_Part1(t *testing.T) {
	d := &day17{input: "target area: x=20..30, y=-10..-5"}
	const expected = "45"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

func TestDay17_Part2(t *testing.T) {
	d := &day17{input: "target area: x=20..30, y=-10..-5"}
	const expected = "112"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}
