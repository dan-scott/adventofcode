package day03

import "testing"

func TestDay03_Part2(t *testing.T) {
	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	d := day03{}
	d.loadLines(lines)
	answer := d.Part2()
	if answer != "230" {
		t.Fatalf("Expected 230 but got %s", answer)
	}
}
