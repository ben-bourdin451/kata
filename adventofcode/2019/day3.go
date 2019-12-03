package adventofcode

import (
	"strconv"
	"strings"
)

type direction struct {
	orientation string
	steps       int
}

func newDirection(s string) direction {
	n, _ := strconv.Atoi(s[1:])
	return direction{string(s[0]), n}
}

func initWire(in string) []direction {
	w := []direction{}
	for _, d := range strings.Split(in, ",") {
		w = append(w, newDirection(d))
	}
	return w
}

func traceWire(r []direction) map[point]int {
	path := make(map[point]int)

	x, y, steps := 0, 0, 0
	for _, d := range r {
		for i := 0; i < d.steps; i++ {
			switch d.orientation {
			case "R":
				x++
				break
			case "L":
				x--
				break
			case "U":
				y--
				break
			case "D":
				y++
				break
			}
			steps++
			if _, visited := path[point{x, y}]; !visited {
				path[point{x, y}] = steps
			}
		}
	}

	return path
}

func findIntersections(w1, w2 map[point]int) []point {
	points := []point{}
	for k := range w1 {
		if _, intersects := w2[k]; intersects {
			points = append(points, k)
		}
	}

	return points
}

func day3Part1(in []string) int {
	w1, w2 := initWire(in[0]), initWire(in[1])

	intersections := findIntersections(traceWire(w1), traceWire(w2))

	origin := point{0, 0}
	closest, _ := origin.closest(intersections)

	return origin.distance(closest)
}

func day3Part2(in []string) int {
	w1, w2 := traceWire(initWire(in[0])), traceWire(initWire(in[1]))
	intersections := findIntersections(w1, w2)

	minSteps := 0
	for _, p := range intersections {
		if sum := w1[p] + w2[p]; minSteps == 0 || sum < minSteps {
			minSteps = sum
		}
	}

	return minSteps
}
