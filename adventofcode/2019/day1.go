package adventofcode

import (
	"math"
	"strconv"
)

func day1Part1(in []string) int {
	sum := 0
	for _, s := range in {
		n, _ := strconv.Atoi(s)
		sum += calculateMass(n)
	}

	return sum
}

func calculateMass(n int) int {
	return int(math.Floor(float64(n)/3)) - 2
}

func day1Part2(in []string) int {
	return totalMass(in)
}

func totalMass(in []string) int {
	total := 0
	for _, s := range in {
		n, _ := strconv.Atoi(s)
		mSum := 0
		for m := calculateMass(n); m > 0; m = calculateMass(m) {
			mSum += m
		}
		total += mSum
	}
	return total
}
