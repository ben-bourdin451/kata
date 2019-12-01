package adventofcode

func day2Part1(in *[]string) int {
	twice, thrice := 0, 0
	for _, s := range *in {
		var m = make(map[rune]int, len(s))
		for _, l := range s {
			if _, ok := m[l]; ok {
				m[l]++
			} else {
				m[l] = 1
			}
		}

		found2, found3 := false, false
		for _, v := range m {
			if v == 2 {
				found2 = true
			} else if v == 3 {
				found3 = true
			}
		}

		if found2 {
			twice++
		}
		if found3 {
			thrice++
		}
	}

	return twice * thrice
}

func day2Part2(in *[]string) string {
	for i, s1 := range *in {
		for j, s2 := range *in {
			if i == j {
				continue
			}

			mismatchCount := 0
			mismatchIndex := 0
			for n := range s1 {
				if s1[n] != s2[n] {
					mismatchCount++
					mismatchIndex = n
				}
			}
			if mismatchCount < 2 {
				return s1[:mismatchIndex] + s1[mismatchIndex+1:]
			}
		}
	}

	return "No matches!"
}
