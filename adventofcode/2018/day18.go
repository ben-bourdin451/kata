package adventofcode

import (
	"fmt"
	"reflect"
	"strings"
)

type lumberCollectionAcre struct {
	x int
	y int
	t string
}

func day18(in []string, minutes int) int {
	area := make([][]string, len(in))
	for y, row := range in {
		area[y] = strings.Split(row, "")
	}

	// fmt.Println("Initial area:")
	// printArea(area)
	savedAreas := [][]lumberCollectionAcre{}
	foundRepeating := false
	patternStart := 500
	for i := 0; i < minutes; i++ {
		// Start checking for repeating patterns after a few generations
		if i >= patternStart && !foundRepeating {
			seenArea := saveArea(area)
			seen, memIndex := checkSavedAreas(savedAreas, seenArea)

			if seen {
				foundRepeating = true
				steps := len(savedAreas) - memIndex

				// Skip ahead
				// -1 because of the for loop
				i += int((minutes-i)/steps)*steps - 1
				continue

			} else {
				savedAreas = append(savedAreas, seenArea)
			}
		}

		nextArea := make([][]string, len(area))
		for y := 0; y < len(area); y++ {
			nextArea[y] = make([]string, len(area[y]))
			for x := 0; x < len(area[y]); x++ {
				_, treeC, lumberC := countSurroundings(area, point{x, y})

				if area[y][x] == "." && treeC >= 3 {
					nextArea[y][x] = "|"

				} else if area[y][x] == "|" && lumberC >= 3 {
					nextArea[y][x] = "#"

				} else if area[y][x] == "#" {
					if lumberC >= 1 && treeC >= 1 {
						nextArea[y][x] = "#"
					} else {
						nextArea[y][x] = "."
					}
				} else {
					nextArea[y][x] = area[y][x]
				}
			}
		}

		area = nextArea
	}

	return woodedAcres(area)
}

func checkSavedAreas(saved [][]lumberCollectionAcre, next []lumberCollectionAcre) (bool, int) {
	for i, area := range saved {
		if compareAcres(area, next) {
			return true, i
		}
	}
	return false, -1
}

func compareAcres(a, b []lumberCollectionAcre) bool {
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func saveArea(area [][]string) []lumberCollectionAcre {
	r := []lumberCollectionAcre{}
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			r = append(r, lumberCollectionAcre{x, y, area[y][x]})
		}
	}
	return r
}

func woodedAcres(area [][]string) int {
	wood, lumber := 0, 0
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			if area[y][x] == "|" {
				wood++
			} else if area[y][x] == "#" {
				lumber++
			}
		}
	}

	return wood * lumber
}

func countSurroundings(area [][]string, p point) (int, int, int) {
	openC, treeC, lumberC := 0, 0, 0

	for y := p.y - 1; y <= p.y+1; y++ {
		if y < 0 || y >= len(area) {
			continue
		}

		for x := p.x - 1; x <= p.x+1; x++ {
			if x < 0 || x >= len(area[y]) || (x == p.x && y == p.y) {
				continue
			}

			switch area[y][x] {
			case ".":
				openC++
			case "|":
				treeC++
			case "#":
				lumberC++
			}
		}
	}

	return openC, treeC, lumberC
}

func printArea(area [][]string) {
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			fmt.Print(area[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println()
}
