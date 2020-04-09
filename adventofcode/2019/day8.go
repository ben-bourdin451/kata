package adventofcode

import (
	"fmt"
	"strconv"
)

type layer struct {
	img   [][]int
	count map[int]int
}

func newLayer(w, h int) *layer {
	img := make([][]int, h)
	for y := range img {
		img[y] = make([]int, w)
	}

	return &layer{img, make(map[int]int)}
}

func initLayers(in string, w, h int) []*layer {
	n := len(in) / (w * h)
	layers := make([]*layer, n)
	for i := range layers {
		layers[i] = newLayer(w, h)
	}

	i := 0
	for _, l := range layers {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				d, _ := strconv.Atoi(string(in[i]))
				i++

				l.img[y][x] = d
				l.count[d]++
			}
		}
	}

	return layers
}

func day8Part1(in string, w, h int) int {
	layers := initLayers(in, w, h)

	fewest := 0
	for i, l := range layers {
		if l.count[0] < layers[fewest].count[0] {
			fewest = i
		}
	}

	return layers[fewest].count[1] * layers[fewest].count[2]
}

func day8Part2(in string, w, h int) {
	layers := initLayers(in, w, h)

	img := make([][]int, h)
	for y := range img {
		img[y] = make([]int, w)
		for x := range img[y] {
			for _, l := range layers {
				if l.img[y][x] != 2 {
					img[y][x] = l.img[y][x]
					if img[y][x] == 1 {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
					break
				}
			}
		}
		fmt.Println()
	}
}
