package adventofcode

import (
	"strconv"
)

func day8Part1(in []string) int {
	_, _, sum := metaSum(in, 0, 0)

	return sum
}

type node struct {
	children []*node
	meta     []int
	val      int
}

func metaSum(in []string, index, sum int) (*node, int, int) {
	c, _ := strconv.Atoi(in[index])
	index++

	m, _ := strconv.Atoi(in[index])
	index++

	children := make([]*node, c)
	for i := range children {
		children[i], index, sum = metaSum(in, index, sum)
	}

	meta := make([]int, m)
	for i := range meta {
		d, _ := strconv.Atoi(in[index])
		index++

		meta[i] = d
		sum += d
	}
	return &node{children, meta, 0}, index, sum
}

func day8Part2(in []string) int {
	root, _ := nodeVal(in, 0)

	return root.val
}

func nodeVal(in []string, index int) (*node, int) {
	c, _ := strconv.Atoi(in[index])
	index++

	m, _ := strconv.Atoi(in[index])
	index++

	children := make([]*node, c)
	for i := range children {
		children[i], index = nodeVal(in, index)
	}

	sum := 0
	meta := make([]int, m)
	for i := range meta {
		d, _ := strconv.Atoi(in[index])
		index++

		meta[i] = d

		if len(children) > 0 && 0 < d && d <= len(children) {
			sum += children[d-1].val // metadata points to child num not index
		} else if len(children) <= 0 {
			sum += d
		}
	}
	return &node{children, meta, sum}, index
}
