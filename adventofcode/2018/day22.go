package adventofcode

import "fmt"

type cave struct {
	depth  int
	target point
	scan   [][]region
	b      *bounds
}

type region struct {
	geo     int
	erosion int
	risk    int
}

type location struct {
	p    point
	tool string
}

type locationDist struct {
	p    point
	tool string
	time int
}

const (
	rocky = iota
	wet
	narrow
)

func day22Part1(depth int, target point) int {
	c := cave{depth, target, nil, nil}
	c.init(0)
	sum := 0
	for y := range c.scan {
		for x := range c.scan[y] {
			sum += c.scan[y][x].risk
		}
	}

	return sum
}

func (c *cave) init(len int) {
	len++ // we start at 0
	c.b = &bounds{-1, -1, c.target.x + len, c.target.y + len}
	c.scan = make([][]region, c.target.y+len)
	for y := 0; y < c.target.y+len; y++ {
		c.scan[y] = make([]region, c.target.x+len)
		for x := range c.scan[y] {
			g := c.getGeoIndex(point{x, y})
			e := c.getErosion(g)
			c.scan[y][x] = region{g, e, e % 3}
		}
	}
}

func (c *cave) getGeoIndex(p point) int {
	if p.x == 0 && p.y == 0 || p.x == c.target.x && p.y == c.target.y {
		return 0

	} else if p.x == 0 {
		return p.y * 48271

	} else if p.y == 0 {
		return p.x * 16807
	}

	return c.scan[p.y][p.x-1].erosion * c.scan[p.y-1][p.x].erosion
}

func (c *cave) getErosion(geo int) int {
	return (geo + c.depth) % 20183
}

func (c *cave) draw() {
	for y := range c.scan {
		for x := range c.scan[y] {
			if x == 0 && y == 0 {
				fmt.Print("M")
			} else if x == c.target.x && y == c.target.y {
				fmt.Print("T")
			} else {
				switch c.scan[y][x].risk {
				case rocky:
					fmt.Print(".")
				case wet:
					fmt.Print("=")
				case narrow:
					fmt.Print("|")
				}
			}
		}
		fmt.Println()
	}
}

func day22Part2(depth int, target point) int {
	c := cave{depth, target, nil, nil}
	c.init(64)

	visited := map[location]locationDist{}
	c.traverse(location{point{0, 0}, "T"}, visited)

	if t, found := visited[location{c.target, "T"}]; found {
		return t.time
	}

	return 0
}

func (c *cave) traverse(start location, visited map[location]locationDist) {
	q := []locationDist{locationDist{start.p, start.tool, 0}}

	for len(q) > 0 {
		// Pop element from Q
		curr := q[0]
		currL := location{curr.p, curr.tool}
		q = q[1:]

		l, seen := visited[currL]
		if seen && l.time <= curr.time {
			continue
		}
		visited[currL] = curr

		// We need to check the possibility of switching tools at current location
		r := c.scan[curr.p.y][curr.p.x]
		q = append(q, locationDist{curr.p, r.otherTool(curr.tool), curr.time + 7})

		// We also need to check how far we can go with current tool
		edges := []point{point{curr.p.x, curr.p.y - 1}, point{curr.p.x - 1, curr.p.y}, point{curr.p.x + 1, curr.p.y}, point{curr.p.x, curr.p.y + 1}}
		for _, p := range edges {
			// Ignore point if it is out of bounds of the cave
			// or if we can't reach it with our current tool
			if !c.b.contains(&p) || !r.canReach(c.scan[p.y][p.x], curr.tool) {
				continue
			}
			q = append(q, locationDist{p, curr.tool, curr.time + 1})
		}
	}
}

// Returns other tool you can switch to in current region
func (r *region) otherTool(t string) string {
	switch r.risk {
	case 0:
		if t == "C" {
			return "T"
		}
		return "C"

	case 1:
		if t == "C" {
			return "N"
		}
		return "C"

	case 2:
		if t == "T" {
			return "N"
		}
		return "T"
	}

	return ""
}

// Return true if we can reach the next region with given tool, otherwise false
func (r *region) canReach(n region, tool string) bool {
	switch tool {
	case "C":
		if n.risk == rocky || n.risk == wet {
			return true
		}
	case "T":
		if n.risk == rocky || n.risk == narrow {
			return true
		}
	case "N":
		if n.risk == wet || n.risk == narrow {
			return true
		}
	}
	return false
}
