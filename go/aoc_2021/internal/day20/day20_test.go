package day20

import "testing"

func TestDay20_Part1(t *testing.T) {
	d := &day20{input}
	d.Open()
	//const expected = "35"
	const expected = "5347"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

func TestDay20_Part2(t *testing.T) {
	d := &day20{input}
	//d.Open()
	const expected = "3351"
	//const expected = "5347"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

var input = []string{
	"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
	"",
	"#..#.",
	"#....",
	"##..#",
	"..#..",
	"..###",
}
