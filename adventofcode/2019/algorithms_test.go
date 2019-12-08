package adventofcode

import "testing"

func TestAlgorithmsDijkstra(t *testing.T) {
	com := planet{"com", nil, nil}
	a := planet{"a", nil, nil}
	b := planet{"b", nil, nil}
	c := planet{"a", nil, nil}
	a.orbitsAround(&com)
	b.orbitsAround(&a)
	c.orbitsAround(&b)

	nodes := []graph{&com, &a, &b, &c}
	dist, _ := dijkstra(nodes, (graph)(&com))

	cDist := dist[(graph)(&c)]
	if cDist != 3 {
		t.Errorf("got %v, want %v", cDist, 3)
	}
}
