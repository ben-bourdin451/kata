package adventofcode

import (
	"reflect"
	"strconv"
	"strings"
)

func day2Part1(in string) int {

	reg := strings.Split(in, ",")
	codes := []int{}
	for _, code := range reg {
		n, _ := strconv.Atoi(code)
		codes = append(codes, n)
	}

	return intcode(codes)
}

func intcode(codes []int) int {
	for i := 0; i < len(codes); {
		switch codes[i] {
		case 1:
			sum := codes[codes[i+1]] + codes[codes[i+2]]
			codes[codes[i+3]] = sum
			break

		case 2:
			sum := codes[codes[i+1]] * codes[codes[i+2]]
			codes[codes[i+3]] = sum
			break
		case 99:
			return codes[0]
		}
		i += 4
	}

	return codes[0]
}

func day2Part2(in string) (int, int) {
	reg := strings.Split(in, ",")
	codes := []int{}
	for _, code := range reg {
		n, _ := strconv.Atoi(code)
		codes = append(codes, n)
	}

	max := 100
	for n := 0; n < max; n++ {
		for v := 0; v < max; v++ {
			cp := make([]int, len(codes))
			reflect.Copy(reflect.ValueOf(cp), reflect.ValueOf(codes))
			cp[1] = n
			cp[2] = v
			res := intcode(cp)
			if res == 19690720 {
				return n, v
			}
		}
	}

	return -1, -1
}
