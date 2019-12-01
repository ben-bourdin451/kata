package adventofcode

func day9Part1(players, limit int) int {
	r := &marble{val: 0}
	m1 := &marble{1, r, r}
	r.next, r.prev = m1, m1

	curr := m1
	score := make([]int, players)
	for i, p := 2, 1; i <= limit; i++ {

		if i%23 == 0 {
			v, n := curr.removePrev(7)
			score[p] = score[p] + i + v
			curr = n
		} else {
			new := &marble{val: i}
			curr.next.add(new)
			curr = new
		}

		p++
		if p >= players {
			p = 0
		}
	}

	best := score[0]
	for _, v := range score {
		if v > best {
			best = v
		}
	}

	return best
}

type marble struct {
	val  int
	prev *marble
	next *marble
}

func (curr *marble) add(m *marble) {
	n := curr.next
	if n != nil {
		m.next = n
		n.prev = m
	}
	curr.next = m
	m.prev = curr
}

func (curr *marble) remove() (int, *marble) {
	p := curr.prev
	n := curr.next
	p.next = n
	n.prev = p
	return curr.val, n
}

func (curr *marble) removePrev(count int) (int, *marble) {
	for i := 0; i < count; i++ {
		curr = curr.prev
	}
	return curr.remove()
}
