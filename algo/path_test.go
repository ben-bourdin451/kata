package algo

import "testing"

func TestAlgorithmsDijkstra(t *testing.T) {
	com := &GNode{Name: "com", Neighbors: []*Edge{}}
	a := &GNode{Name: "a", Neighbors: []*Edge{}}
	b := &GNode{Name: "b", Neighbors: []*Edge{}}
	c := &GNode{Name: "c", Neighbors: []*Edge{}}

	com.Neighbors = append(com.Neighbors, &Edge{a, 1})
	a.Neighbors = append(a.Neighbors, &Edge{b, 1})
	b.Neighbors = append(b.Neighbors, &Edge{c, 1})

	nodes := []Graph{com, a, b, c}
	dist, _ := dijkstra(nodes, com)

	cDist := dist[c]
	if cDist != 3 {
		t.Errorf("got %v, want %v", cDist, 3)
	}
}
