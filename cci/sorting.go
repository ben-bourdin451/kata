package cci

// Bubble sort
func bubbleSort(arr []int) {
	for {
		swapped := false

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}
