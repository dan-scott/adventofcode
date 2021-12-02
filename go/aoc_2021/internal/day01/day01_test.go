package day01

import "testing"

func TestDay01_Part1(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part1()
	if answer != "1696" {
		t.Fatalf("Expected 1696 but got %s", answer)
	}
	d.Close()
}

func TestDay01_Part2(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part2()
	if answer != "1737" {
		t.Fatalf("Expected 1737 but got %s", answer)
	}
	d.Close()
}

func BenchmarkDay01_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay01_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
