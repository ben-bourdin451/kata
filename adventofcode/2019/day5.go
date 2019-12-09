package adventofcode

func day5(mem string, input int) int {
	memory := initCodes(mem)

	_, output := intcode(memory, []int{input})

	for _, out := range output {
		if out != 0 {
			return out
		}
	}

	return -1
}
