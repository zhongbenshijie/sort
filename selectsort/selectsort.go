package bubblesort

func SelectSort(intSlice []int) {
	for i := 0; i < len(intSlice)-1; i++ {
		for j := i + 1; j < len(intSlice); j++ {
			if intSlice[i] > intSlice[j] {
				t := intSlice[i]
				intSlice[i] = intSlice[j]
				intSlice[j] = t
			}
		}
	}
}
