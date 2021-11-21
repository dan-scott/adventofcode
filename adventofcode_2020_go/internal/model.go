package internal

type Day interface {
	Open()
	Close()
	Part1() string
	Part2() string
}
