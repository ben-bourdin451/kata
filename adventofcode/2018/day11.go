package adventofcode

import (
	"math"
)

func day11Part1(serial int) *point {

	// 1-300 for 3x3 grid --> 0-297 --> 1-298
	best, topLeft := 0, &point{0, 0}
	for y := 1; y <= 298; y++ {
		for x := 1; x <= 298; x++ {
			sum := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					sum += fuelCellPower(x+i, y+j, serial)
				}
			}

			if sum > best {
				best = sum
				topLeft.x = x
				topLeft.y = y
			}
		}
	}

	return topLeft
}

func fuelCellPower(x, y, serial int) int {
	rack := x + 10
	power := rack*y + serial
	power *= rack
	power = int(math.Trunc(float64(power/100))) % 10
	return power - 5
}

func day11Part2(serial int) (*point, int) {
	// 1-300 for 3x3 grid --> 0-297 --> 1-298
	best, topLeft, size := 0, &point{0, 0}, 0
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			for n := 0; n <= 300-x && n <= 300-y; n++ {
				sum := 0
				for i := 0; i < n; i++ {
					for j := 0; j < n; j++ {
						sum += fuelCellPower(x+i, y+j, serial)
					}
				}

				if sum > best {
					best = sum
					size = n
					topLeft.x = x
					topLeft.y = y
				}
			}
		}
	}

	return topLeft, size
}
