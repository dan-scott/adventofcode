package day07

import "testing"

func TestDay07_Part1(t *testing.T) {
	d := &day07{}
	d.load("16,1,2,0,4,2,7,1,2,14")
	answer := d.Part1()
	if answer != "37" {
		t.Fatalf("Expected 37 but got %s", answer)
	}
}
