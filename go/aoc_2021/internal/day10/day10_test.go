package day10

import "testing"

func TestDay10_Part2(t *testing.T) {
	d := &day10{}
	d.loadLines([]string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	})
	const expected = "288957"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}
