package adventofcode

import (
	"reflect"
)

func day2Part1(in string) int {
	codes := initCodes(in)
	mem, _ := intcode(codes, []int{0})
	return mem[0]
}

func day2Part2(in string) int {
	codes := initCodes(in)

	max := 100
	for n := 0; n < max; n++ {
		for v := 0; v < max; v++ {
			cp := make([]int, len(codes))
			reflect.Copy(reflect.ValueOf(cp), reflect.ValueOf(codes))
			cp[1] = n
			cp[2] = v
			mem, _ := intcode(cp, []int{0})
			if mem[0] == 19690720 {
				return n*100 + v
			}
		}
	}

	return 0
}
