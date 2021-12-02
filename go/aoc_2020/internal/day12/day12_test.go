package day12

import "testing"

func TestMD(t *testing.T) {
	lines := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	d := &day12{}
	d.loadLines(lines)
	md := d.Part1()
	if md != "25" {
		t.Fatalf("Expected 25 but got %s", md)
	}
}

func TestWPMD(t *testing.T) {
	lines := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	d := &day12{}
	d.loadLines(lines)
	md := d.Part2()
	if md != "286" {
		t.Fatalf("Expected 286 but got %s", md)
	}
}
