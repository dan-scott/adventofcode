package day04

import "testing"

func TestDay04_Part1(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part1()
	if answer != "46920" {
		t.Fatalf("Expected 46920 but got %s", answer)
	}
}

func TestDay04_Part2(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part2()
	if answer != "12635" {
		t.Fatalf("Expected 12635 but got %s", answer)
	}
}

func BenchmarkDay04_getInitialBoards(b *testing.B) {
	d := &day04{}
	d.Open()
	for i := 0; i < b.N; i++ {
		d.getInitialBoards()
	}
}

func BenchmarkDay04_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay04_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
