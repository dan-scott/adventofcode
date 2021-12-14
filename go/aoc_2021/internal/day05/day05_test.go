package day05

import "testing"

func TestDay05_Part1(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part1()
	if answer != "5442" {
		t.Fatalf("Expected 5442 but got %s", answer)
	}
}

func TestDay05_Part2(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part2()
	if answer != "19571" {
		t.Fatalf("Expected 19571 but got %s", answer)
	}
}

func BenchmarkDay05_loadLines(b *testing.B) {
	d := &day05{}
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.getLines()
	}
}

func BenchmarkDay05_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay05_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
