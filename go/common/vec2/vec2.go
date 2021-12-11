package vec2

import "fmt"

type Vec2 struct {
	X, Y int
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%d %d)", v.X, v.Y)
}

func Of(x, y int) Vec2 {
	return Vec2{x, y}
}

func OfIndex(index, run int) Vec2 {
	return Vec2{index % run, index / run}
}

func OfInt64(x, y int64) Vec2 {
	return Vec2{int(x), int(y)}
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Of(v.X+o.X, v.Y+o.Y)
}

func (v Vec2) Index(run int) int {
	return v.Y*run + v.X
}
