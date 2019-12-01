package adventofcode

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type nanobot struct {
	p point3d
	r int
}

type point3d struct {
	x, y, z int
}

type bounds3d struct {
	minX, minY, minZ int
	maxX, maxY, maxZ int
}

func (b *bounds3d) update(p point3d) {
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

	if p.z < b.minZ {
		b.minZ = p.z
	} else if b.maxZ < p.z {
		b.maxZ = p.z
	}
}

func newNanobot(s string) nanobot {
	regx := regexp.MustCompile(`^pos=<\s*(?P<x>[0-9-]+),\s*(?P<y>[0-9-]+),\s*(?P<z>[0-9-]+)>, r=(?P<r>[0-9-]+)`)
	m := regx.FindStringSubmatch(s)
	x, y, z, r := 0, 0, 0, 0
	for i, n := range regx.SubexpNames() {
		e, _ := strconv.Atoi(m[i])
		switch n {
		case "x":
			x = e
		case "y":
			y = e
		case "z":
			z = e
		case "r":
			r = e
		}
	}
	return nanobot{point3d{x, y, z}, r}
}

func (p *point3d) distance(other point3d) int {
	return int(math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y)) + math.Abs(float64(p.z-other.z)))
}

func (b *nanobot) hasInRange(p point3d) bool {
	return p.distance(b.p) <= b.r
}

func (b *nanobot) rangeIntersects(other nanobot) bool {
	return !(b.p.distance(other.p) > b.r+other.r)
}

func (b *nanobot) instersectingBounds() bounds3d {
	return bounds3d{0, 0, 0, 0, 0, 0}
}

func day23Part1(in []string) int {
	strongest, bots, _ := initBots(in)

	count := 0
	for _, b := range bots {
		if strongest.hasInRange(b.p) {
			count++
		}
	}
	return count
}

func initBots(in []string) (nanobot, []nanobot, bounds3d) {
	bots := []nanobot{}
	var strongest *nanobot
	max := 10000
	b := bounds3d{max, max, max, 0, 0, 0}
	for _, s := range in {
		bot := newNanobot(s)
		b.update(bot.p)
		if strongest == nil || bot.r > strongest.r {
			strongest = &bot
		}
		bots = append(bots, bot)
	}
	return *strongest, bots, b
}

func day23Part2(in []string) int {
	return solution1(in)
}

func solution3(in []string) int {
	_, bots, _ := initBots(in)

	ib := intersectionBounds(bots)

	bestc, bestPos := 0, []point3d{}
	for _, bound := range ib {
		c, pos := findMostInRange(bound, bots)
		if bestc == 0 || c < bestc {
			bestc = c
			bestPos = pos
		} else if c == bestc {
			bestPos = append(bestPos, pos...)
		}
	}

	closest := findClosestDistance(bestPos)
	return closest
}

func intersectionBounds(bots []nanobot) []bounds3d {
	q := bots
	r := []bounds3d{}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, b := range q {
			if b.rangeIntersects(curr) {
				r = append(r, b.instersectingBounds())
			}
		}
	}
	return r
}

// Returns the position in the given bounds that has the most bots in range
func findMostInRange(r bounds3d, bots []nanobot) (int, []point3d) {
	bestc, bestPos := 0, []point3d{}
	for z := r.minZ; z < r.maxZ; z++ {
		for y := r.minY; y < r.maxY; y++ {
			for x := r.minX; x < r.maxX; x++ {
				p := point3d{x, y, z}
				count := 0
				for _, b := range bots {
					if b.hasInRange(p) {
						count++
					}
				}

				if count > bestc {
					bestc = count
					bestPos = []point3d{p}
				} else if count == bestc {
					bestPos = append(bestPos, p)
				}
			}
		}
	}

	return bestc, bestPos
}

func findClosestDistance(points []point3d) int {
	closest := 0
	for _, p := range points {
		if d := p.distance(point3d{0, 0, 0}); closest == 0 || d < closest {
			closest = d
		}
	}
	return closest
}

func solution2(in []string) int {
	_, bots, _ := initBots(in)

	bestb := mostIntersect(bots)
	fmt.Println("found best:", bestb)
	bestBounds := bounds3d{
		minX: bestb.p.x,
		minY: bestb.p.y,
		minZ: bestb.p.z,
		maxX: bestb.p.x + bestb.r,
		maxY: bestb.p.y + bestb.r,
		maxZ: bestb.p.z + bestb.r,
	}

	_, bestPos := findMostInRange(bestBounds, bots)

	closest := findClosestDistance(bestPos)
	return closest
}

func mostIntersect(bots []nanobot) nanobot {
	q := bots
	best, bestb := 0, q[0]
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		icount := 0
		for _, b := range bots {
			if b.rangeIntersects(curr) {
				icount++
			}
		}

		if icount > best {
			best = icount
			bestb = curr
		}
	}
	return bestb
}

func solution1(in []string) int {
	_, bots, b := initBots(in)
	fmt.Println(b)

	_, bestPos := findMostInRange(b, bots)

	closest := findClosestDistance(bestPos)
	return closest
}
