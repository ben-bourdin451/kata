package main

import (
	"flag"
	"fmt"
	"math"
)

type point struct {
	x, y int
}

const (
	left  = 0
	up    = 1
	right = 2
	down  = 3
)

var input = flag.Int("input", 4, "Input size for the spiral")

func main() {
	flag.Parse()
	in := *input

	max := int(math.Pow(float64(in), float64(2)))
	grid := make([][]int, in)
	for i := range grid {
		grid[i] = make([]int, in)
	}

	s := &spiral{grid, max, 1}
	p := &point{0, 0}
	steps := in
	for i := 0; s.count <= s.max && steps > 0; i++ {
		s.crawl(p, steps, right)
		steps--

		p.move(down)
		s.crawl(p, steps, down)

		p.move(left)
		s.crawl(p, steps, left)
		steps--

		p.move(up)
		s.crawl(p, steps, up)

		p.move(right)
	}

	printGrid(grid)
}

type spiral struct {
	grid  [][]int
	max   int
	count int
}

func printGrid(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%3d", grid[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println()
}

func (s *spiral) crawl(p *point, steps, direction int) {
	for i := 0; i < steps && s.count <= s.max; i++ {
		s.grid[p.y][p.x] = s.count
		s.count++

		if i < steps-1 {
			p.move(direction)
		}
	}
}

func (p *point) move(direction int) {
	switch direction {
	case up:
		p.y--
		break
	case right:
		p.x++
		break
	case down:
		p.y++
		break
	case left:
		p.x--
		break
	}
}
