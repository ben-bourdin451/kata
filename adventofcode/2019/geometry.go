package adventofcode

import (
	"math"
)

const (
	left  = 0
	up    = 1
	right = 2
	down  = 3
)

type point struct {
	x, y int
}

func (p *point) closest(coor []point) (point, bool) {
	shortest := 1000
	dupe := false
	var closest point
	for _, c := range coor {
		d := c.distance(*p)
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

func (p *point) distanceSum(coor []point) int {
	sum := 0
	for _, c := range coor {
		sum += p.distance(c)
	}

	return sum
}

func (p *point) distance(p2 point) int {
	return int(math.Abs(float64(p.x-p2.x)) + math.Abs(float64(p.y-p2.y)))
}

func (p *point) nextPoint(direction, dist int) point {
	switch direction {
	case left:
		return point{p.x - dist, p.y}
	case up:
		return point{p.x, p.y - dist}
	case right:
		return point{p.x + dist, p.y}
	case down:
		return point{p.x, p.y + dist}
	}

	return point{-1, -1}
}

type bounds struct {
	minX int
	minY int
	maxX int
	maxY int
}

func newBounds() bounds {
	return bounds{int(math.Inf(-1)), int(math.Inf(-1)), int(math.Inf(0)), int(math.Inf(0))}
}

func (b *bounds) contains(p point) bool {
	return b.minX < p.x && b.minY < p.y && p.x < b.maxX && p.y < b.maxY
}

func (b *bounds) update(p point) {
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
