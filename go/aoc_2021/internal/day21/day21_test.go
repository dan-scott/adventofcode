package day21

import (
	"testing"
)

func TestDay21_Part1(t *testing.T) {
	d := &day21{inputs: []string{
		"Player 1 starting position: 4",
		"Player 2 starting position: 8",
	}}
	const expected = "739785"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}
func TestDay21_Part2(t *testing.T) {
	d := &day21{inputs: []string{
		"Player 1 starting position: 4",
		"Player 2 starting position: 8",
	}}
	const expected = "444356092776315"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}
