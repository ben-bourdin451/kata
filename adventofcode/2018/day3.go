package adventofcode

import (
	"regexp"
	"strconv"
)

type claim struct {
	id   int
	x, y int
	w, h int
}

func parseClaim(c string) claim {
	r := regexp.MustCompile(`^#(?P<id>\d{1,5}) @ (?P<x>\d{1,5}),(?P<y>\d{1,5}): (?P<w>\d{1,5})x(?P<h>\d{1,5})`)
	match := r.FindStringSubmatch(c)

	id, x, y, w, h := 0, 0, 0, 0, 0
	for i, name := range r.SubexpNames() {
		n, _ := strconv.Atoi(match[i])
		switch name {
		case "id":
			id = n
		case "x":
			x = n
		case "y":
			y = n
		case "w":
			w = n
		case "h":
			h = n
		}
	}
	return claim{id, x, y, w, h}
}

func processClaim(c claim, fabric [][]int) {
	endY := c.y + c.h
	endX := c.x + c.w
	for y := c.y; y < endY; y++ {
		for x := c.x; x < endX; x++ {
			fabric[y][x]++
		}
	}
}

func day3Part1(in *[]string) int {
	fabric := make([][]int, 1000)
	for i := range fabric {
		fabric[i] = make([]int, 1000)
	}

	for _, claim := range *in {
		processClaim(parseClaim(claim), fabric)
	}

	overlap := 0
	for y, row := range fabric {
		for x := range row {
			if fabric[y][x] > 1 {
				overlap++
			}
		}
	}

	return overlap
}

func imprint(c claim, fabric [][][]int) {
	endY := c.y + c.h
	endX := c.x + c.w
	for y := c.y; y < endY; y++ {
		for x := c.x; x < endX; x++ {
			fabric[y][x] = append(fabric[y][x], c.id)
		}
	}
}

func day3Part2(in *[]string) int {
	fabric := make([][][]int, 1000)
	for i := range fabric {
		fabric[i] = make([][]int, 1000)
	}

	for _, claim := range *in {
		imprint(parseClaim(claim), fabric)
	}

	uniqueIds := map[int]struct{}{}
	overlappingIds := map[int]struct{}{}
	for y, row := range fabric {
		for x := range row {
			if len(fabric[y][x]) > 1 {
				// This part of the fabric has overlap
				// Add all ids to our known overlapping claims
				for _, id := range fabric[y][x] {
					overlappingIds[id] = struct{}{}
					if _, ok := uniqueIds[id]; ok {
						// Also delete any previously stored unique claims
						delete(uniqueIds, id)
					}
				}
			} else if len(fabric[y][x]) == 1 {
				// This part of the fabric has only 1 claim
				id := fabric[y][x][0]
				if _, ok := overlappingIds[id]; !ok {
					// If not in our known overlapping claims
					// Then Add to unique claims
					uniqueIds[id] = struct{}{}
				}
			}
		}
	}

	var result int
	for k := range uniqueIds {
		result = k
	}

	return result
}
