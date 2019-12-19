package adventofcode

import (
	"reflect"
)

func day7Part1(in string) int {
	memory := initCodes(in)

	signal := 0
	phases := []int{0, 1, 2, 3, 4}
	for _, perm := range permutations(phases, []int{}) {
		cpMem := make([]int, len(memory))
		reflect.Copy(reflect.ValueOf(cpMem), reflect.ValueOf(memory))

		out := 0
		for _, p := range perm {
			input := []int{p, out}
			mem, output := intcode(cpMem, input)
			cpMem, out = mem, output[0]
		}
		if out > signal {
			signal = out
		}
	}

	return signal
}

func day7Part2(in string) int {
	memory := initCodes(in)

	signal := 0
	phases := []int{9, 7, 8, 5, 6}
	// for _, perm := range permutations(phases, []int{}) {
	cpMem := make([]int, len(memory))
	reflect.Copy(reflect.ValueOf(cpMem), reflect.ValueOf(memory))

	out := 0
	for done := false; !done; {
		for _, p := range phases {
			input := []int{p, out}
			mem, output := intcode(cpMem, input)
			if len(output) == 0 {
				done = true
				break
			}
			cpMem, out = mem, output[0]
		}
	}
	if out > signal {
		signal = out
	}
	// }

	return signal
}
