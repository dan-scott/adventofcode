package day19

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
)

type day19 struct {
	lines []string
}

type report struct {
	id      string
	beacons []vec3
	pairIds []vec3
	distMap map[vec3][]vec3
	offset  vec3
	solved  bool
}

func (r *report) matches(solved *report) bool {
	matchCt := 0
	matched := make([]vec3, 0)
	for _, k := range r.pairIds {
		if _, ok := solved.distMap[k]; ok {
			matchCt++
			matched = append(matched, k)
		}
	}
	if matchCt < 66 {
		return false
	}
	basePair := solved.distMap[matched[0]]
	newPair := r.distMap[matched[0]]
	rotIdx := 0
	diff := v(0, 0, 0)
	found := false
	for i, rotate := range rotations {
		d1, match := checkDiffs(basePair[0], basePair[1], rotate(newPair[0]), rotate(newPair[1]))
		if !match {
			continue
		}
		for _, k := range matched[1:] {
			bp := solved.distMap[k]
			np := r.distMap[k]
			d2, m2 := checkDiffs(bp[0], bp[1], rotate(np[0]), rotate(np[1]))
			if !m2 || d1 != d2 {
				continue
			}
		}
		found = true
		rotIdx = i
		diff = d1
		break
	}
	if !found {
		return false
	}
	rot := rotations[rotIdx]
	r.setOrientation(rot, diff)
	return true
}

func (r *report) setOrientation(rotation vecRotation, offset vec3) {
	r.solved = true
	r.offset = offset
	for i, b := range r.beacons {
		r.beacons[i] = rotation(b).add(offset)
	}
	for k := range r.distMap {
		r.distMap[k][0] = rotation(r.distMap[k][0]).add(offset)
		r.distMap[k][1] = rotation(r.distMap[k][1]).add(offset)
	}
}

func checkDiffs(b0, b1, n0, n1 vec3) (vec3, bool) {
	d1 := b0.sub(n0)
	if b1 == n1.add(d1) {
		return d1, true
	}
	d2 := b0.sub(n1)
	if b1 == n0.add(d2) {
		return d2, true
	}
	return d2, false
}

func (d *day19) Open() {
	d.lines = inputs.LinesAsString(2021, 19)
}

func (d *day19) Close() {
	d.lines = nil
}

func (d *day19) Part1() string {
	reports := d.parseReports()
	solve(reports)
	beaconMap := make(map[vec3]interface{})
	for _, r := range reports {
		for _, b := range r.beacons {
			beaconMap[b] = nil
		}
	}
	return fmt.Sprint(len(beaconMap))

}

func solve(reports []*report) {
	solved := make(map[string]*report, 4)
	reports[0].solved = true
	solved[reports[0].id] = reports[0]
	for len(solved) < len(reports) {
		for _, r := range reports {
			if r.solved {
				continue
			}
			for _, s := range solved {
				if r.matches(s) {
					solved[r.id] = r
					break
				}
			}
		}
	}
}

func (d *day19) Part2() string {
	reports := d.parseReports()
	solve(reports)
	max := 0
	for i, r := range reports[:len(reports)-1] {
		for _, r2 := range reports[i+1:] {
			md := r.offset.md(r2.offset)
			if md > max {
				max = md
			}
		}
	}
	return fmt.Sprint(max)
}

func (d *day19) parseReports() []*report {
	reports := make([]*report, 0)
	idx := 0
	for idx < len(d.lines) {
		id2 := idx
		for id2 < len(d.lines) && d.lines[id2] != "" {
			id2++
		}
		reports = append(reports, parseReport(d.lines[idx:id2]))
		idx = id2 + 1
	}
	return reports
}

func parseReport(lines []string) *report {
	id := lines[0]
	beacons := make([]vec3, len(lines)-1)
	for i, l := range lines[1:] {
		beacons[i] = parse(l)
	}
	distMap := make(map[vec3][]vec3)
	ids := make([]vec3, 0)
	for i := 0; i < len(beacons)-1; i++ {
		for j := i + 1; j < len(beacons); j++ {
			k := beacons[i].diffId(beacons[j])
			distMap[k] = []vec3{beacons[i], beacons[j]}
			ids = append(ids, k)
		}
	}
	return &report{
		id:      id,
		beacons: beacons,
		distMap: distMap,
		pairIds: ids,
	}
}

func New() runner.Day {
	return &day19{}
}
