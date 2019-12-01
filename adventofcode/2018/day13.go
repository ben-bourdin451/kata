package adventofcode

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	left  = 0
	up    = 1
	right = 2
	down  = 3
)

func day13Part1(in []string) *point {
	track, carts := initTrack(in)
	for {
		// printTrack(track, carts)
		for _, c := range carts {
			n := c.nextPos()
			if i := containsCart(n, carts); i >= 0 {
				return n
			}

			c.move(track[n.y][n.x])
		}
	}
}

func initTrack(in []string) ([][]string, []*cart) {
	track := make([][]string, len(in))
	carts := []*cart{}
	for y, row := range in {
		track[y] = strings.Split(row, "")
		for x := range track[y] {
			if d, ok := isCart(track[y][x]); ok {
				carts = append(carts, &cart{&point{x, y}, d, 0, false})
				if d == up || d == down {
					track[y][x] = "|"
				} else {
					track[y][x] = "-"
				}
			}
		}
	}

	return track, carts
}

func printTrack(track [][]string, carts []*cart) {
	for y := range track {
		for x := range track[y] {
			if i := containsCart(&point{x, y}, carts); i >= 0 {
				switch carts[i].direction {
				case left:
					fmt.Print("<")
				case up:
					fmt.Print("^")
				case right:
					fmt.Print(">")
				case down:
					fmt.Print("v")
				}
			} else {
				fmt.Print(track[y][x])
			}
		}
		fmt.Println()
	}
	time.Sleep(1000 * time.Millisecond)
}

func containsCart(p *point, carts []*cart) int {
	for i, c := range carts {
		if !c.broken && reflect.DeepEqual(c.pos, p) {
			return i
		}
	}

	return -1
}

func isCart(s string) (int, bool) {
	switch s {
	case "<":
		return left, true
	case "^":
		return up, true
	case ">":
		return right, true
	case "v":
		return down, true
	}

	return -1, false
}

type cart struct {
	pos       *point
	direction int
	turnCount int
	broken    bool
}

func (c *cart) move(next string) {
	c.pos = c.nextPos()
	c.changeDirection(next)
}

func (c *cart) nextPos() *point {
	switch c.direction {
	case left:
		return &point{c.pos.x - 1, c.pos.y}
	case up:
		return &point{c.pos.x, c.pos.y - 1}
	case right:
		return &point{c.pos.x + 1, c.pos.y}
	case down:
		return &point{c.pos.x, c.pos.y + 1}
	}

	return &point{-1, -1}
}

func (c *cart) turnLeft() {
	switch c.direction {
	case left:
		c.direction = down
	case up:
		c.direction = left
	case right:
		c.direction = up
	case down:
		c.direction = right
	}
}

func (c *cart) turnRight() {
	switch c.direction {
	case left:
		c.direction = up
	case up:
		c.direction = right
	case right:
		c.direction = down
	case down:
		c.direction = left
	}
}

func (c *cart) changeDirection(n string) {
	switch n {
	case `\`:
		if c.direction == down || c.direction == up {
			c.turnLeft()
		} else {
			c.turnRight()
		}
	case `/`:
		if c.direction == down || c.direction == up {
			c.turnRight()
		} else {
			c.turnLeft()
		}
	case `+`:
		switch c.turnCount % 3 {
		case 0:
			c.turnLeft()
		case 2:
			c.turnRight()
		}
		c.turnCount++
	}
}

func unbrokenCarts(carts []*cart) int {
	s := 0
	for _, c := range carts {
		if !c.broken {
			s++
		}
	}
	return s
}

func day13Part2(in []string) *point {
	track, carts := initTrack(in)
	for unbrokenCarts(carts) > 1 {
		// printTrack(track, carts)
		for _, c := range carts {
			if c.broken {
				continue
			}

			n := c.nextPos()
			if j := containsCart(n, carts); j >= 0 {
				c.broken = true
				carts[j].broken = true
			} else {
				c.move(track[n.y][n.x])
			}
		}
	}

	for _, c := range carts {
		if !c.broken {
			return c.pos
		}
	}

	return &point{0, 0}
}
