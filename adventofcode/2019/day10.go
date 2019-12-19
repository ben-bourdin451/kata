package adventofcode

import (
	"strings"
)

func day10Part1(in []string) int {
	space := make([][]string, len(in))
	b := bounds{-1, -1, len(in[0]), len(in)}

	asteroids := make(map[point]int)
	for y := range in {
		space[y] = strings.Split(in[y], "")
		for x := range space[y] {
			if space[y][x] == "#" {
				asteroids[point{x, y}] = 0
			}
		}
	}

	best := 0
	for p := range asteroids {
		c := detectAsteroids(space, asteroids, b, p)
		if c > best {
			best = c
		}
		asteroids[p] = c
	}

	return best
}

func detectAsteroids(space [][]string, asteroids map[point]int, b bounds, origin point) int {
	// for k, v := range asteroids {
	// 	if reflect.DeepEqual(k, origin) {
	// 		continue
	// 	}

	// }

	return 0
}

func day10Part2(in []string) int {
	return 0
}
