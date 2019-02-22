package bubblesort

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	intSlice := []int{2, 401, 58, 5, 1, 6, 29, 189}
	SelectSort(intSlice)

	t.Log(intSlice)
}
