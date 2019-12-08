package adventofcode

import (
	"strconv"
)

type criteria func(i int) bool

func part1(i int) bool { return i >= 2 }
func part2(i int) bool { return i == 2 }

func day4(min, max int, c criteria) int {
	count := 0

	for i := min; i <= max; i++ {
		if passwordMatch(i, c) {
			count++
		}
	}

	return count
}

func passwordMatch(p int, adjacent criteria) bool {
	prev := -1
	digits := make(map[int]int)
	for _, d := range strconv.Itoa(p) {
		digit, _ := strconv.Atoi(string(d))
		digits[digit]++

		if prev > digit {
			return false
		}

		prev = digit
	}

	r := false
	for _, v := range digits {
		if adjacent(v) {
			r = true
		}
	}

	return r
}
