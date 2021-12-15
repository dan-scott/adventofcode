package day15

import "testing"

func TestDay15_Part1(t *testing.T) {
	d := &day15{
		lines: []string{
			"1163751742",
			"1381373672",
			"2136511328",
			"3694931569",
			"7463417111",
			"1319128137",
			"1359912421",
			"3125421639",
			"1293138521",
			"2311944581",
		},
	}
	//d.Open()
	const expected = "40"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
	d.Close()
}

func TestDay15_Part2(t *testing.T) {
	d := &day15{
		lines: []string{
			"1163751742",
			"1381373672",
			"2136511328",
			"3694931569",
			"7463417111",
			"1319128137",
			"1359912421",
			"3125421639",
			"1293138521",
			"2311944581",
		},
	}
	//d.Open()
	const expected = "315"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
	d.Close()
}

func BenchmarkDay15_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay15_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
