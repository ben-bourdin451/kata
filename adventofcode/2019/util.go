package adventofcode

import (
	"bufio"
	"math"
	"os"
)

func readStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var in []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	return in, nil
}

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

type bounds struct {
	minX int
	minY int
	maxX int
	maxY int
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
