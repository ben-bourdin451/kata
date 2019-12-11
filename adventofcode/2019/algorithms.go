package adventofcode

import (
	"math"
	"reflect"
)

type graph interface {
	// returns the neighbors of the node alongside the distance to them
	getNeighbors() ([]graph, []float64)
}

// returns shortest distances & paths to each node from the source
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
		unvisited = append(unvisited[:i], unvisited[i+1:]...)

		neighbors, nDist := u.getNeighbors()
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

func closest(unvisited []graph, dist map[graph]float64) int {
	lowest, index := math.Inf(0), 0
	unique := make(map[graph]float64)
	for i, u := range unvisited {
		if d, ok := dist[u]; ok && d < lowest {
			unique[u] = d
			lowest, index = d, i
		}
	}

	return index
}

func permutations(in []int, prefix []int) [][]int {
	if len(in) == 0 {
		return [][]int{prefix}
	}

	perms := [][]int{}
	for i := range in {
		cp := make([]int, len(in))
		reflect.Copy(reflect.ValueOf(cp), reflect.ValueOf(in))
		np := append(prefix, cp[i])
		rem := append(cp[:i], cp[i+1:]...)
		perms = append(perms, permutations(rem, np)...)
	}
	return perms
}
