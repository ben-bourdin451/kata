package adventofcode

import (
	"reflect"
)

func day2Part1(argv string) int64 {
	mem := initCodes(argv)
	input := make(chan int64, 1)
	defer close(input)
	input <- 0
	intcode(mem, input, nil)

	return mem[0]
}

func day2Part2(argv string) int64 {
	codes := initCodes(argv)
	input := make(chan int64, 1)
	defer close(input)
	input <- 0

	max := 100
	for n := 0; n < max; n++ {
		for v := 0; v < max; v++ {
			cp := make([]int64, len(codes))
			reflect.Copy(reflect.ValueOf(cp), reflect.ValueOf(codes))
			cp[1] = int64(n)
			cp[2] = int64(v)
			intcode(cp, input, nil)
			if cp[0] == 19690720 {
				return int64(n*100 + v)
			}
		}
	}
	return 0
}
