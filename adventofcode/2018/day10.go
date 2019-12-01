package adventofcode

import (
	"fmt"
	"regexp"
	"strconv"
)

func day10(input []string, minArea int) int {
	b := &bounds{10000, 10000, 0, 0}
	sky := make([]*light, 0, len(input))
	for _, in := range input {
		l := parseLight(in)
		sky = append(sky, l)
		b.update(l.p)
	}

	sec := 0
	for b.area() > minArea {
		b = &bounds{10000, 10000, 0, 0}
		for _, l := range sky {
			l.move()
			b.update(l.p)
		}
		drawSky(sky, b, minArea)
		sec++
	}

	return sec
}

func drawSky(sky []*light, b *bounds, minArea int) {
	if b.area() <= minArea {
		for y := b.minY; y <= b.maxY; y++ {
			for x := b.minX; x <= b.maxX; x++ {
				if isInSky(sky, &point{x, y}) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}

func isInSky(sky []*light, p *point) bool {
	for _, l := range sky {
		if l.p.x == p.x && l.p.y == p.y {
			return true
		}
	}
	return false
}

type light struct {
	p *point
	v *point
}

func (l *light) move() {
	l.p.x = l.p.x + l.v.x
	l.p.y = l.p.y + l.v.y
}

func parseLight(s string) *light {
	r := regexp.MustCompile(`^position=<\s*(?P<x>[0-9-]+),\s*(?P<y>[0-9-]+)> velocity=<\s*(?P<vx>[0-9-]+),\s*(?P<vy>[0-9-]+)>`)
	m := r.FindStringSubmatch(s)
	x, y, vx, vy := 0, 0, 0, 0
	for i, n := range r.SubexpNames() {
		e, _ := strconv.Atoi(m[i])
		switch n {
		case "x":
			x = e
		case "y":
			y = e
		case "vx":
			vx = e
		case "vy":
			vy = e
		}
	}
	return &light{&point{x, y}, &point{vx, vy}}
}
