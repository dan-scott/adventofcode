package day22

import (
	"fmt"
	"math"
)

type vec3 struct {
	x, y, z int
}

type componentId = int

const (
	cX componentId = iota
	cY
	cZ
)

func (vc vec3) String() string {
	return fmt.Sprintf("v(%d, %d, %d)", vc.x, vc.y, vc.z)
}

func v(x, y, z int) vec3 {
	return vec3{x, y, z}
}

func (vc vec3) getCmp(i int) int {
	switch i {
	case cX:
		return vc.x
	case cY:
		return vc.y
	case cZ:
		return vc.z
	}
	panic("invalid component")
}

func (vc vec3) setCmp(i, val int) vec3 {
	switch i {
	case cX:
		return v(val, vc.y, vc.z)
	case cY:
		return v(vc.x, val, vc.z)
	case cZ:
		return v(vc.x, vc.y, val)
	}
	panic("invalid component")
}

func (vc vec3) abs() vec3 {
	return vec3{
		x: int(math.Abs(float64(vc.x))),
		y: int(math.Abs(float64(vc.y))),
		z: int(math.Abs(float64(vc.z))),
	}
}

func (vc vec3) prod() int {
	return vc.x * vc.y * vc.z
}

func (vc vec3) sub(o vec3) vec3 {
	return vec3{
		vc.x - o.x,
		vc.y - o.y,
		vc.z - o.z,
	}
}

func (vc vec3) add(o vec3) vec3 {
	return vec3{
		vc.x + o.x,
		vc.y + o.y,
		vc.z + o.z,
	}
}
