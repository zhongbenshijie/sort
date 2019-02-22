package bubblesort

// 冒泡排序与选择排序的区别在于,冒泡是相邻比较
func BubbleSort(intSlice []int) {
	for i := 0; i < len(intSlice)-1; i++ {
		for j := 0; j < len(intSlice)-1-i; j++ {
			if intSlice[j] > intSlice[j+1] {
				t := intSlice[j]
				intSlice[j] = intSlice[j+1]
				intSlice[j+1] = t
			}
		}
	}
}
