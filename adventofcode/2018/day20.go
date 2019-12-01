package adventofcode

type room struct {
	distance int
	p        point // Location
}

func newRoom(rooms map[point]*room, prev *room, d string) *room {
	var p point
	switch d {
	case "N":
		p = point{prev.p.x, prev.p.y - 1}
	case "E":
		p = point{prev.p.x + 1, prev.p.y}
	case "S":
		p = point{prev.p.x, prev.p.y + 1}
	case "W":
		p = point{prev.p.x - 1, prev.p.y}
	}

	// Check if we've already visited this room
	if r, ok := rooms[p]; ok {
		return r
	}

	rooms[p] = &room{prev.distance + 1, p}
	return rooms[p]
}

func discover(in string) map[point]*room {
	rooms := map[point]*room{}
	stack := []*room{}
	prev := &room{0, point{0, 0}}
	for _, c := range in {
		switch c {
		case '(':
			stack = append(stack, prev)
		case '|':
			prev = stack[len(stack)-1]
		case ')':
			stack = stack[:len(stack)-1]
		case 'N', 'S', 'W', 'E':
			nr := newRoom(rooms, prev, string(c))
			prev = nr
		}
	}

	return rooms
}

func day20Part1(in string) int {
	rooms := discover(in)

	max := 0
	for _, v := range rooms {
		if v.distance > max {
			max = v.distance
		}
	}

	return max
}

func day20Part2(in string, distance int) int {
	rooms := discover(in)

	count := 0
	for _, v := range rooms {
		if v.distance >= distance {
			count++
		}
	}

	return count
}
