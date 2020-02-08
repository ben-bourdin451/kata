package adventofcode

func day5(mem string, argv int64) int64 {
	memory := initCodes(mem)
	in, out := make(chan int64, 1), make(chan int64)
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
