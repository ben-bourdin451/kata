package adventofcode

import (
	"math"
	"strconv"
	"strings"
)

func day6Part1(in []string) int {

	points := []*point{}
	areaCount := make(map[*point]int, len(in))
	b := &bounds{1000, 1000, 0, 0}
	for _, s := range in {
		p := pointFromString(s)
		points = append(points, p)
		areaCount[p] = 0
		b.update(p)
	}

	blacklist := map[*point]struct{}{}
	for y := b.minY; y <= b.maxY; y++ {
		for x := b.minX; x <= b.maxX; x++ {
			p := &point{x, y}

			if c, dupe := p.closest(points); dupe {
				continue

			} else if !b.contains(p) {
				blacklist[c] = struct{}{}

			} else {
				areaCount[c] = areaCount[c] + 1
			}
		}
	}

	best := 0
	for k, v := range areaCount {
		if _, bl := blacklist[k]; v > best && !bl {
			best = v
		}
	}
	return best
}

type point struct {
	x, y int
}

func (p *point) closest(coor []*point) (*point, bool) {
	shortest := 1000
	dupe := false
	var closest *point
	for _, c := range coor {
		d := c.distance(p)
		if d < shortest {
			shortest = d
			closest = c
			dupe = false
		} else if d == shortest {
			dupe = true
		}
	}

	return closest, dupe
}

func (p *point) distanceSum(coor []*point) int {
	sum := 0
	for _, c := range coor {
		sum += p.distance(c)
	}

	return sum
}

func (p *point) distance(p2 *point) int {
	return int(math.Abs(float64(p.x-p2.x)) + math.Abs(float64(p.y-p2.y)))
}

func pointFromString(str string) *point {
	s := strings.Split(str, ",")
	x, _ := strconv.Atoi(strings.TrimSpace(s[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(s[1]))

	return &point{x, y}
}

type bounds struct {
	minX int
	minY int
	maxX int
	maxY int
}

func (b *bounds) contains(p *point) bool {
	return b.minX < p.x && b.minY < p.y && p.x < b.maxX && p.y < b.maxY
}

func (b *bounds) update(p *point) {
	if p.x < b.minX {
		b.minX = p.x
	} else if b.maxX < p.x {
		b.maxX = p.x
	}

	if p.y < b.minY {
		b.minY = p.y
	} else if b.maxY < p.y {
		b.maxY = p.y
	}
}

func (b *bounds) area() int {
	return (b.maxX - b.minX) * (b.maxY - b.minY)
}

func day6Part2(in []string, limit int) int {
	points := []*point{}
	b := &bounds{1000, 1000, 0, 0}
	for _, s := range in {
		p := pointFromString(s)
		points = append(points, p)
		b.update(p)
	}

	area := 0
	for y := b.minY; y <= b.maxY; y++ {
		for x := b.minX; x <= b.maxX; x++ {
			p := &point{x, y}

			if p.distanceSum(points) < limit {
				area++
			}
		}
	}
	return area
}
