package bubblesort

// 选择排序可以以最大为选择目的, 也可以以最小为选择目的
func SelectSort(intSlice []int) {
	for i := 0; i < len(intSlice)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(intSlice); j++ {
			if intSlice[minIndex] > intSlice[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			t := intSlice[i]
			intSlice[i] = intSlice[minIndex]
			intSlice[minIndex] = t
		}
	}
}
