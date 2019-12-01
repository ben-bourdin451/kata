package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

type soil struct {
	b       *bounds
	clay    map[point]struct{}
	water   map[point]struct{}
	springs map[point]struct{}
}

func day17Part1(in []string) int {
	s := initSoil(in)
	s.fall(&point{500, 0}, false)

	return len(s.water)
}

func initSoil(in []string) *soil {
	b := &bounds{2100, 2100, 0, 0}
	clay := make(map[point]struct{})
	water := make(map[point]struct{})
	springs := make(map[point]struct{})

	for _, line := range in {
		for _, p := range parseClay(line) {
			b.update(&p)
			clay[p] = struct{}{}
		}
	}

	return &soil{b, clay, water, springs}
}

func (s *soil) fall(source *point, drains bool) {
	curr := *source

	for {
		below := point{curr.x, curr.y + 1}
		if _, ok := s.clay[below]; ok {
			s.water[curr] = struct{}{}
			break
		} else if curr.y >= s.b.maxY {
			return
		}
		if !drains {
			s.water[below] = struct{}{}
		}
		curr = below
	}

	s.fill(&curr, drains)

	return
}

func (s *soil) fill(source *point, drains bool) {
	curr := *source
	lS := s.spread(&curr, left)
	rS := s.spread(&curr, right)

	// While water can't go down
	// Then keep spreading water and going up
	for lS == nil && rS == nil {
		curr.y--
		s.water[curr] = struct{}{}
		lS = s.spread(&curr, left)
		rS = s.spread(&curr, right)
	}

	hwl := false
	if lS != nil {
		s.newSpring(lS, curr, drains)
		lls := point{lS.x - 1, lS.y}
		_, hwl = s.water[lls]
	}

	hwr := false
	if rS != nil {
		s.newSpring(rS, curr, drains)
		rrs := point{rS.x + 1, rS.y}
		_, hwr = s.water[rrs]
	}

	if drains && !(hwl || hwr) {
		s.drain(curr, left)
		s.drain(curr, right)
	}

	return
}

func (s *soil) drain(curr point, direction int) {
	delete(s.water, curr)
	for {
		_, c := s.clay[curr]
		_, sp := s.springs[curr]

		if direction == left {
			curr.x--
		} else {
			curr.x++
		}

		if c || sp {
			break
		}

		delete(s.water, curr)
	}
}

func (s *soil) newSpring(source *point, curr point, drains bool) {
	if _, ok := s.springs[*source]; !ok {
		s.springs[*source] = struct{}{}
		s.fall(source, drains)
	}
}

func (s *soil) spread(source *point, direction int) *point {
	curr := source
	for {

		next := point{curr.x + 1, curr.y}
		nextNext := point{next.x + 1, next.y}
		if direction == left {
			next = point{curr.x - 1, curr.y}
			nextNext = point{next.x - 1, next.y}
		}
		belowNext := point{next.x, next.y + 1}

		_, cn := s.clay[next]
		_, sn := s.springs[next]
		_, wnn := s.water[nextNext]
		_, cbn := s.clay[belowNext]
		_, wbn := s.water[belowNext]

		// If clay or spring filled up with water
		if cn || (sn && wnn) {
			return nil // water keeps going up
		} else if sn {
			// Otherwise if we hit another unfilled spring we need to go down it
			return &next
		}
		// else we hit nothing and need to add water
		s.water[next] = struct{}{}

		// If there is no water or clay below the next one then we need to fall (create new spring)
		if !cbn && !wbn {
			return &next
		}

		curr = &next
	}
}

func (s *soil) printSoil() {
	for y := s.b.minY; y <= s.b.maxY; y++ {
		for x := s.b.minX; x <= s.b.maxX; x++ {
			curr := point{x, y}
			if _, c := s.clay[curr]; c {
				fmt.Print("#")
			} else if _, spring := s.springs[curr]; spring {
				fmt.Print("+")
			} else if _, w := s.water[curr]; w {
				fmt.Print("^")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|", y)
	}
	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println()
}

func parseClay(s string) []point {
	p := []point{}
	parts := strings.Split(s, ", ")

	left, _ := strconv.Atoi(strings.Split(parts[0], "=")[1])
	rightParts := strings.Split(strings.Split(parts[1], "=")[1], "..")
	rightMin, _ := strconv.Atoi(rightParts[0])
	rightMax, _ := strconv.Atoi(rightParts[1])

	for i := rightMin; i <= rightMax; i++ {
		if strings.Contains(parts[0], "x") {
			p = append(p, point{left, i})
		} else {
			p = append(p, point{i, left})
		}
	}

	return p
}

func day17Part2(in []string) int {
	s := initSoil(in)

	s.fall(&point{500, 0}, true)
	s.printSoil()

	return len(s.water)
}
