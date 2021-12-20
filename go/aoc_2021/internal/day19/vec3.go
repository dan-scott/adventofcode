package day19

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type vec3 struct {
	x, y, z int
}

type vecRotation = func(vi vec3) vec3

var rotations = []vecRotation{
	func(vi vec3) vec3 { return v(-vi.x, -vi.y, vi.z) },
	func(vi vec3) vec3 { return v(-vi.x, -vi.z, -vi.y) },
	func(vi vec3) vec3 { return v(-vi.x, vi.y, -vi.z) },
	func(vi vec3) vec3 { return v(-vi.x, vi.z, vi.y) },
	func(vi vec3) vec3 { return v(-vi.y, -vi.x, -vi.z) },
	func(vi vec3) vec3 { return v(-vi.y, -vi.z, vi.x) },
	func(vi vec3) vec3 { return v(-vi.y, vi.x, vi.z) },
	func(vi vec3) vec3 { return v(-vi.y, vi.z, -vi.x) },
	func(vi vec3) vec3 { return v(-vi.z, -vi.x, vi.y) },
	func(vi vec3) vec3 { return v(-vi.z, -vi.y, -vi.x) },
	func(vi vec3) vec3 { return v(-vi.z, vi.x, -vi.y) },
	func(vi vec3) vec3 { return v(-vi.z, vi.y, vi.x) },
	func(vi vec3) vec3 { return v(vi.x, -vi.y, -vi.z) },
	func(vi vec3) vec3 { return v(vi.x, -vi.z, vi.y) },
	func(vi vec3) vec3 { return v(vi.x, vi.y, vi.z) },
	func(vi vec3) vec3 { return v(vi.x, vi.z, -vi.y) },
	func(vi vec3) vec3 { return v(vi.y, -vi.x, vi.z) },
	func(vi vec3) vec3 { return v(vi.y, -vi.z, -vi.x) },
	func(vi vec3) vec3 { return v(vi.y, vi.x, -vi.z) },
	func(vi vec3) vec3 { return v(vi.y, vi.z, vi.x) },
	func(vi vec3) vec3 { return v(vi.z, -vi.x, -vi.y) },
	func(vi vec3) vec3 { return v(vi.z, -vi.y, vi.x) },
	func(vi vec3) vec3 { return v(vi.z, vi.x, vi.y) },
	func(vi vec3) vec3 { return v(vi.z, vi.y, -vi.x) },
}

func (v vec3) String() string {
	return fmt.Sprintf("v(%d, %d, %d)", v.x, v.y, v.z)
}

func v(x, y, z int) vec3 {
	return vec3{x, y, z}
}

func parse(line string) vec3 {
	pts := strings.Split(line, ",")
	x, _ := strconv.Atoi(pts[0])
	y, _ := strconv.Atoi(pts[1])
	z, _ := strconv.Atoi(pts[2])
	return v(x, y, z)
}

func (v vec3) abs() vec3 {
	return vec3{
		x: int(math.Abs(float64(v.x))),
		y: int(math.Abs(float64(v.y))),
		z: int(math.Abs(float64(v.z))),
	}
}

func (v vec3) diffId(o vec3) vec3 {
	d := v.sub(o).abs()
	return vec3{d.sum(), 0, d.max()}
}

func (v vec3) sum() int {
	return v.x + v.y + v.z
}

func (v vec3) max() int {
	if v.x > v.y && v.x >= v.z {
		return v.x
	}
	if v.y >= v.x && v.y >= v.z {
		return v.y
	}
	return v.z
}

func (v vec3) md(o vec3) int {
	return v.sub(o).abs().sum()
}

func (v vec3) sub(o vec3) vec3 {
	return vec3{
		v.x - o.x,
		v.y - o.y,
		v.z - o.z,
	}
}

func (v vec3) add(o vec3) vec3 {
	return vec3{
		v.x + o.x,
		v.y + o.y,
		v.z + o.z,
	}
}

func (v vec3) rot(a vec3) vec3 {
	return v.rotX(a.x).rotY(a.y).rotZ(a.z)
}

func (v vec3) rotX(n int) vec3 {
	nv := vec3{v.x, v.y, v.z}
	for i := 0; i < n; i++ {
		nv.y, nv.z = -nv.z, nv.y
	}
	return nv
}

func (v vec3) rotZ(n int) vec3 {
	nv := vec3{v.x, v.y, v.z}
	for i := 0; i < n; i++ {
		nv.x, nv.y = nv.y, -nv.x
	}
	return nv
}

func (v vec3) rotY(n int) vec3 {
	nv := vec3{v.x, v.y, v.z}
	for i := 0; i < n; i++ {
		nv.x, nv.z = -nv.z, nv.x
	}
	return nv
}
