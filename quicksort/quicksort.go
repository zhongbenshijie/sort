package quicksort

func QuickSort(intSlice []int, start int, end int) {
	if start < end {
		i := start
		j := end
		t := intSlice[i]

		for i < j {
			for i < j {
				if intSlice[j] >= t {
					j = j - 1
				} else {
					break
				}
			}

			if i < j {
				intSlice[i] = intSlice[j]
				i = i + 1
			}

			for i < j {
				if intSlice[i] <= t {
					i = i + 1
				} else {
					break
				}
			}

			if i < j {
				intSlice[j] = intSlice[i]
				j = j - 1
			}

			if i == j {
				break
			}
		}

		intSlice[i] = t
		QuickSort(intSlice, start, i-1)
		QuickSort(intSlice, i+1, end)
	}
}
