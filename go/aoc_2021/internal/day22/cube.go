package day22

import (
	"fmt"
)

type cube struct {
	min, max vec3
}

func (c cube) String() string {
	return fmt.Sprintf("[%s %s]", c.min, c.max)
}

func (c cube) volume() int {
	return c.max.sub(c.min).abs().add(v(1, 1, 1)).prod()
}

func (c cube) overlappingCube(o cube) (cube, bool) {
	if !c.intersects(o) {
		return cube{}, false
	}
	return cube{
		min: c.min.maxVec3(o.min),
		max: c.max.minVec3(o.max),
	}, true
}

func (c cube) intersects(o cube) bool {
	return c.cmpTouches(o, cX) && c.cmpTouches(o, cY) && c.cmpTouches(o, cZ)
}

func (c cube) cmpTouches(o cube, i componentId) bool {
	cMin, cMax := c.cmpRange(i)
	oMin, oMax := o.cmpRange(i)

	if oMax < cMin || oMin > cMax {
		return false
	}

	return true
}

func (c cube) cmpRange(i componentId) (int, int) {
	return c.min.getCmp(i), c.max.getCmp(i)
}

func (c cube) containsPoint(v vec3) (a, x, y, z bool) {
	x = v.x >= c.min.x && v.x <= c.max.x
	y = v.y >= c.min.y && v.y <= c.max.y
	z = v.z >= c.min.z && v.z <= c.max.z
	a = x && y && z
	return
}

func (c cube) contains(o cube) bool {
	ma, _, _, _ := c.containsPoint(o.min)
	mb, _, _, _ := c.containsPoint(o.max)
	return ma && mb
}
