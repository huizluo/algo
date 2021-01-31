package _6sorts

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	MergeSort(arr)
	t.Log(arr)

	arr = []int{5, 3, 2, 4, 1}
	MergeSort(arr)
	t.Log(arr)
}