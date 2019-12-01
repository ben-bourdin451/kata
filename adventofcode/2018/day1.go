package adventofcode

func day1Part1(in *[]int) int {
	var s int
	for _, n := range *in {
		s += n
	}
	return s
}

func day1Part2(in *[]int) int {
	s := 0
	m := map[int]struct{}{0: {}}
	for {
		for _, n := range *in {
			s += n
			if _, ok := m[s]; ok {
				return s
			}
			m[s] = struct{}{}
		}
	}
}
