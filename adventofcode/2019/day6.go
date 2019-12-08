package adventofcode

import (
	"strings"
)

type planet struct {
	name      string
	orbits    []*planet
	inOrbitOf []*planet
}

func (p planet) String() string {
	return p.name
}

func (p *planet) orbitsAround(center *planet) {
	p.orbits = append(p.orbits, center)
	center.inOrbitOf = append(center.inOrbitOf, p)
}

func (p *planet) countOrbits() int {
	sum := 0
	for _, orbits := range p.orbits {
		sum += orbits.countOrbits() + 1
	}
	return sum
}

func (p *planet) getNeighbors() ([]graph, []float64) {
	n, d := []graph{}, []float64{}

	for _, p := range p.orbits {
		n = append(n, p)
		d = append(d, 1)
	}
	for _, p := range p.inOrbitOf {
		n = append(n, p)
		d = append(d, 1)
	}

	return n, d
}

func day6Part1(in []string) int {
	planets := initPlanets(in)
	total := 0
	for _, v := range planets {
		total += v.countOrbits()
	}
	return total
}

func initPlanets(in []string) map[string]*planet {
	planets := make(map[string]*planet)
	for _, orbit := range in {
		rel := strings.Split(orbit, ")")
		n1, n2 := rel[0], rel[1]

		if _, exists := planets[n1]; !exists {
			planets[n1] = &planet{n1, nil, nil}
		}
		if _, exists := planets[n2]; !exists {
			planets[n2] = &planet{n2, nil, nil}
		}

		planets[n2].orbitsAround(planets[n1])
	}

	return planets
}

func day6Part2(in []string) int {
	planets := initPlanets(in)

	nodes := []graph{}
	var source graph
	var target graph
	for _, p := range planets {
		nodes = append(nodes, (graph)(p))
		if p.name == "YOU" {
			source = (graph)(p)
		} else if p.name == "SAN" {
			target = (graph)(p)
		}
	}
	dist, _ := dijkstra(nodes, source)

	return int(dist[target]) - 2
}
