package adventofcode

func day5(mem string, argv int) int {
	memory := initCodes(mem)
	in, out := make(chan int, 1), make(chan int)
	defer close(out)
	go intcode(memory, in, out)
	in <- argv
	close(in)

	for n := range out {
		if n != 0 {
			return n
		}
	}

	return -1
}
