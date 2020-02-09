package adventofcode

func day9Part1(argv string) []int64 {
	memory := initCodes(argv)
	in, out := make(chan int64, 1), make(chan int64)
	go intcode(memory, in, out)
	in <- 1
	output := []int64{}
	for o := range out {
		output = append(output, o)
	}
	close(in)

	return output
}

func day9Part2(argv string) []int64 {
	memory := initCodes(argv)
	in, out := make(chan int64, 1), make(chan int64)
	go intcode(memory, in, out)
	in <- 2
	output := []int64{}
	for o := range out {
		output = append(output, o)
	}
	close(in)

	return output
}
