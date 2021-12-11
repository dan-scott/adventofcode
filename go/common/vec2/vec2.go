package vec2

import "fmt"

type Vec2 struct {
	X, Y int64
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%d %d)", v.X, v.Y)
}

func Of(x, y int64) Vec2 {
	return Vec2{x, y}
}

func OfInt(x, y int) Vec2 {
	return Vec2{int64(x), int64(y)}
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Of(v.X+o.X, v.Y+o.Y)
}
