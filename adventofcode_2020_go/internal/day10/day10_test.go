package day10

import "testing"

func TestDistribution(t *testing.T) {
	lines := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	d := &day10{}
	d.loadLines(lines)

	dist := d.Part1()
	if dist != "220" {
		t.Fatalf("Expected 220 but got %s", dist)
	}
}
