package bubblesort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	intSlice := []int{2, 401, 58, 5, 1, 6, 29, 189}
	BubbleSort(intSlice)

	t.Log(intSlice)
}
