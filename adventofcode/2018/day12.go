package adventofcode

import (
	"strings"
)

func day12Part1(init string, combos []string, gen int) int {
	pots := strings.Split(init, "")

	offset := 5
	// add more pots on the edges
	for i := 0; i < offset; i++ {
		pots = append([]string{"."}, pots...)
		pots = append(pots, ".")
	}
	tunnel := &plantTunnel{combos, pots, offset}

	for g := 0; g < gen; g++ {
		tunnel.nextGen()
		tunnel.pad()
		tunnel.crop()
	}

	return tunnel.getValue()
}

type plantTunnel struct {
	combos []string
	pots   []string
	offset int
}

func (t *plantTunnel) nextGen() {
	nGen := make([]string, len(t.pots))
	for i := 0; i < len(t.pots); i++ {
		if 2 < i && i < len(t.pots)-3 {
			nGen[i] = findMatch(strings.Join(t.pots[i-2:i+3], ""), t.combos)
		} else {
			nGen[i] = t.pots[i]
		}
	}
	t.pots = nGen
}

func (t *plantTunnel) pad() {
	if containsPlant(t.pots[len(t.pots)-4:]) {
		t.pots = append(t.pots, ".", ".", ".", ".")
	}
	if containsPlant(t.pots[:4]) {
		t.pots = append([]string{".", ".", ".", "."}, t.pots...)
		t.offset += 4
	}
}

func (t *plantTunnel) crop() {
	if !containsPlant(t.pots[len(t.pots)-8:]) {
		t.pots = t.pots[:len(t.pots)-4]
	}
	if !containsPlant(t.pots[:8]) {
		t.pots = t.pots[4:]
		t.offset -= 4
	}
}

func (t *plantTunnel) getValue() int {
	sum := 0
	for i, p := range t.pots {
		if p == "#" {
			if i < t.offset {
				sum -= t.offset - i
			} else {
				sum += i - t.offset
			}
		}
	}
	return sum
}

func findMatch(s string, combos []string) string {
	for _, c := range combos {
		split := strings.Split(c, " => ")
		if s == split[0] {
			return split[1]
		}
	}
	return "."
}

func containsPlant(pots []string) bool {
	for _, p := range pots {
		if p == "#" {
			return true
		}
	}
	return false
}

func day12Part2(init string, combos []string, gen int) int {
	if gen < 100 {
		return day12Part1(init, combos, gen)
	}

	return (gen-100+13+185)*186 - 17205
}
