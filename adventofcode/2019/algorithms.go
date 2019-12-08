package adventofcode

import (
	"fmt"
	"math"
)

type vertex struct {
	dist float64
	prev *string
}

type graph interface {
	// returns the neighbors of the node alongside the distance to them
	getNeighbors() ([]graph, []float64)
}

// returns maps of distance & previous nodes from the source
func dijkstra(nodes []graph, source graph) (map[graph]float64, map[graph]graph) {

	unvisited := nodes
	dist := make(map[graph]float64)
	prev := make(map[graph]graph)

	for _, v := range nodes {
		dist[v] = math.Inf(0)
		prev[v] = nil
	}
	dist[source] = 0

	for len(unvisited) > 0 {
		i := closest(unvisited, dist)
		u := unvisited[i]
		// fmt.Println("looking at", u, dist[u])
		unvisited = append(unvisited[:i], unvisited[i+1:]...)

		neighbors, nDist := u.getNeighbors()
		for i, v := range neighbors {
			// fmt.Println("neighbor", v, dist[u]+nDist[i])
			alt := dist[u] + nDist[i]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return dist, prev
}

func closest(unvisited []graph, dist map[graph]float64) int {
	lowest, index := math.Inf(0), 0
	unique := make(map[graph]float64)
	for i, u := range unvisited {
		if d, ok := dist[u]; ok && d < lowest {
			unique[u] = d
			lowest, index = d, i
		}
	}

	if lowest >= math.Inf(0) {
		fmt.Println("could not find closest match")
		fmt.Println(unvisited)
		fmt.Println(unique)
	}

	return index
}
