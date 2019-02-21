package quicksort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	intSlice := []int{2, 401, 58, 5, 1, 6, 29, 189}
	QuickSort(intSlice, 0, len(intSlice)-1)

	t.Log(intSlice)
}
