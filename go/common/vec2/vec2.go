package vec2

type Vec2 struct {
	X, Y int64
}

func Of(x, y int64) Vec2 {
	return Vec2{x, y}
}
