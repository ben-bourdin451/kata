package algo

import (
	"math"
)

// returns shortest distances & paths to each node from the source
func dijkstra(nodes []Graph, source Graph) (map[Graph]float64, map[Graph]Graph) {

	unvisited := nodes
	dist := make(map[Graph]float64)
	prev := make(map[Graph]Graph)

	for _, v := range nodes {
		dist[v] = math.Inf(0)
		prev[v] = nil
	}
	dist[source] = 0

	for len(unvisited) > 0 {
		i := closest(unvisited, dist)
		u := unvisited[i]
		unvisited = append(unvisited[:i], unvisited[i+1:]...)

		neighbors, nDist := u.GetNeighbors()
		for i, v := range neighbors {
			alt := dist[u] + nDist[i]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return dist, prev
}

func closest(unvisited []Graph, dist map[Graph]float64) int {
	lowest, index := math.Inf(0), 0
	unique := make(map[Graph]float64)
	for i, u := range unvisited {
		if d, ok := dist[u]; ok && d < lowest {
			unique[u] = d
			lowest, index = d, i
		}
	}

	return index
}
