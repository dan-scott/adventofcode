package day03

import "testing"

func TestDay03_Part1(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part1()
	if answer != "3687446" {
		t.Fatalf("Expected 3687446 but got %s", answer)
	}
}

func TestDay03_Part2(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part2()
	if answer != "4406844" {
		t.Fatalf("Expected 4406844 but got %s", answer)
	}
}

func TestDay03_Part2_sample(t *testing.T) {
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

func BenchmarkDay03_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay03_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
