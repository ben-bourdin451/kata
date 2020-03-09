package cci

// Bubble sort
func bubbleSort(q []int) []int {
	r := make([]int, len(q))
	copy(r, q)
	if len(q) <= 1 {
		return r
	}

	for {
		swapped := false

		for i := 0; i < len(r)-1; i++ {
			if r[i] > r[i+1] {
				temp := r[i]
				r[i] = r[i+1]
				r[i+1] = temp
				swapped = true
			}
		}

		if !swapped {
			return r
		}
	}
}
