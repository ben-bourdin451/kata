package adventofcode

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
	"time"
)

func day15Part1(in []string) int {
	cave, elves, goblins := initCave(in)

	turns, _ := fight(cave, elves, goblins, 3, false)

	sum := 0
	for _, u := range append(elves, goblins...) {
		if !u.isDead() {
			sum += u.hp
		}
	}

	return turns * sum
}

func initCave(in []string) ([][]string, []*unit, []*unit) {
	cave := make([][]string, len(in))
	elves, goblins := []*unit{}, []*unit{}
	for y, row := range in {
		cave[y] = strings.Split(row, "")

		for x, tile := range cave[y] {
			if tile == "G" {
				g := &unit{&point{x, y}, 200, "G"}
				goblins = append(goblins, g)
			} else if tile == "E" {
				e := &unit{&point{x, y}, 200, "E"}
				elves = append(elves, e)
			}
		}
	}

	return cave, elves, goblins
}

// Returns the number of turns played and if we managed to complete the game
func fight(cave [][]string, e []*unit, g []*unit, elfPower int, stopOnDeath bool) (int, bool) {
	turns := 0
	fullGame := true
	for aliveUnitsCount(e) > 0 && aliveUnitsCount(g) > 0 {
		fullRound, elfDies := playRound(cave, e, g, elfPower)
		if elfDies && stopOnDeath {
			fullGame = false
		}

		if !fullRound {
			break
		}

		turns++
	}

	return turns, fullGame
}

// Returns false if combat finished mid-round and true if an elf died
func playRound(cave [][]string, e []*unit, g []*unit, p int) (bool, bool) {
	elfDies := false
	units := append([]*unit{}, e...)
	units = append(units, g...)
	sortUnits(units)
	for _, u := range units {
		// If fighting ends mid turn we haven't completed a "full round"
		if aliveUnitsCount(e) <= 0 || aliveUnitsCount(g) <= 0 {
			return false, elfDies
		}

		if u.isDead() {
			continue
		}
		move, d := u.findMoveTowardsUnit(cave, e, g)

		if d > 0 {
			cave[move.y][move.x] = cave[u.pos.y][u.pos.x]
			cave[u.pos.y][u.pos.x] = "."
			u.pos = move
		}

		if e := u.findLowestHpEnemyOnAdjacentPoints(e, g); e != nil {
			if e.kind == "G" {
				u.attack(e, p)
			} else {
				u.attack(e, 3)
			}

			if e.isDead() {
				cave[e.pos.y][e.pos.x] = "."
				if e.kind == "E" {
					elfDies = true
				}
			}
		}
	}
	return true, elfDies
}

func printCave(turn int, cave [][]string, e []*unit, g []*unit) {
	fmt.Println("Turn ", turn, " elves left: ", aliveUnitsCount(e), " goblins left: ", aliveUnitsCount(g))
	for y := range cave {
		unitsInRow := []*unit{}
		for x := range cave[y] {
			if u, ok := containsUnit(&point{x, y}, e); ok {
				fmt.Print("E")
				unitsInRow = append(unitsInRow, u)
			} else if u, ok := containsUnit(&point{x, y}, g); ok {
				fmt.Print("G")
				unitsInRow = append(unitsInRow, u)
			} else {
				fmt.Print(cave[y][x])
			}
		}

		for _, u := range unitsInRow {
			fmt.Print(" ", u.printUnit())
		}
		fmt.Println()
	}
	time.Sleep(600 * time.Millisecond)
}

func containsUnit(p *point, units []*unit) (*unit, bool) {
	for _, u := range units {
		if !u.isDead() && reflect.DeepEqual(u.pos, p) {
			return u, true
		}
	}

	return nil, false
}

func sortUnits(u []*unit) {
	sort.Slice(u, func(i, j int) bool {
		if u[i].pos.y == u[j].pos.y {
			return u[i].pos.x < u[j].pos.x
		}
		return u[i].pos.y < u[j].pos.y
	})
}

func sortPoints(p []*point) {
	sort.Slice(p, func(i, j int) bool {
		if p[i].y == p[j].y {
			return p[i].x < p[j].x
		}
		return p[i].y < p[j].y
	})
}

func aliveUnitsCount(units []*unit) int {
	s := 0
	for _, u := range units {
		if !u.isDead() {
			s++
		}
	}
	return s
}

type unit struct {
	pos  *point
	hp   int
	kind string
}

func (u *unit) printUnit() string {
	return fmt.Sprintf("%v(%v)", u.kind, u.hp)
}

func (u *unit) isDead() bool {
	return u.hp <= 0
}

func (u *unit) attack(e *unit, p int) {
	if !e.isDead() {
		e.hp -= p
	}
}

func getReachableAdjacentPoints(cave [][]string, pos *point) []*point {
	up, left, right, down := &point{pos.x, pos.y - 1}, &point{pos.x - 1, pos.y}, &point{pos.x + 1, pos.y}, &point{pos.x, pos.y + 1}
	p := []*point{}
	if cave[up.y][up.x] == "." {
		p = append(p, up)
	}
	if cave[left.y][left.x] == "." {
		p = append(p, left)
	}
	if cave[right.y][right.x] == "." {
		p = append(p, right)
	}
	if cave[down.y][down.x] == "." {
		p = append(p, down)
	}

	return p
}

func (u *unit) findLowestHpEnemyOnAdjacentPoints(e []*unit, g []*unit) *unit {
	if u.kind == "G" {
		return u.findLowestHpUnitOnAdjacentPoints(e)
	}

	return u.findLowestHpUnitOnAdjacentPoints(g)
}

func (u *unit) findLowestHpUnitOnAdjacentPoints(units []*unit) *unit {
	lowest := 201
	var low *unit
	for _, e := range units {
		if !e.isDead() && e.hp < lowest && u.pos.distance(e.pos) == 1 {
			lowest = e.hp
			low = e
		}
	}

	return low
}

func (u *unit) findMoveTowardsUnit(cave [][]string, e []*unit, g []*unit) (*point, int) {
	if u.kind == "G" {
		return u.findMoveTowardsEnemy(cave, e)
	}

	return u.findMoveTowardsEnemy(cave, g)
}

func (u *unit) findMoveTowardsEnemy(cave [][]string, enemies []*unit) (*point, int) {
	// Check we can move
	adjacent := getReachableAdjacentPoints(cave, u.pos)
	if len(adjacent) <= 0 {
		return u.pos, 0
	}

	// If we already have an enemy close to us we don't need to move
	targets := []*point{}
	for _, e := range enemies {
		if !e.isDead() && u.pos.distance(e.pos) == 1 {
			return u.pos, 0
		}

		t := getReachableAdjacentPoints(cave, e.pos)
		if e.isDead() || len(t) <= 0 {
			continue
		} else {
			targets = append(targets, t...)
		}
	}

	// Sort targets in reading order
	sortPoints(targets)

	// Find reachable points
	reachable := u.findReachablePoints(cave)

	// Check which one is closest
	cPoint := u.pos
	cDistance := -1
	for _, t := range targets {
		if v, ok := reachable[*t]; ok && (v.distance < cDistance || cDistance < 0) {
			cDistance = v.distance
			cPoint = v.origin
		}
	}

	return cPoint, cDistance
}

type visitedPoint struct {
	origin   *point
	distance int
}

func (u *unit) findReachablePoints(cave [][]string) map[point]*visitedPoint {
	reachable := make(map[point]*visitedPoint)

	// Find all reachable points
	adjacent := getReachableAdjacentPoints(cave, u.pos)
	for _, p := range adjacent {
		reachable[*p] = &visitedPoint{p, 1}
		lookAhead(cave, p, p, 1, reachable)
	}

	return reachable
}

func lookAhead(cave [][]string, origin *point, curr *point, distance int, reachable map[point]*visitedPoint) {
	adjacent := getReachableAdjacentPoints(cave, curr)
	if len(adjacent) > 0 {
		for _, p := range adjacent {
			if r, known := reachable[*p]; !known || distance+1 < r.distance {
				reachable[*p] = &visitedPoint{origin, distance + 1}
				lookAhead(cave, origin, p, distance+1, reachable)
			}
		}
	}
}

func day15Part2(in []string) int {
	cave, elves, goblins := initCave(in)

	elfRatio := math.Ceil(float64(len(goblins)) / float64(len(elves)))
	if elfRatio < 0 {
		elfRatio = 1
	}

	turns, fullGame := 0, false
	elfPower := int(3 * elfRatio)
	for !fullGame {
		turns, fullGame = fight(cave, elves, goblins, elfPower, true)

		// Reset if we couldn't complete a full game, i.e. an elf died
		if !fullGame {
			elfPower++
			cave, elves, goblins = initCave(in)
			turns = 0
		}
	}

	sum := 0
	for _, u := range append(elves, goblins...) {
		if !u.isDead() {
			sum += u.hp
		}
	}

	return turns * sum
}
