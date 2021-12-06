package day06

import "testing"

func TestDay06_Part1(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part1()
	if answer != "372300" {
		t.Fatalf("Expected 372300 but got %s", answer)
	}
}

func TestDay06_Part2(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part2()
	if answer != "1675781200288" {
		t.Fatalf("Expected 1675781200288 but got %s", answer)
	}
}

func BenchmarkDay06_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay06_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
