package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	a := []int{25, 4, 12, -83, 7, 9, -10, 55, 3, 1}
	MergeSort(a, 0, len(a)-1)
	assert.Equal(t, []int{-83, -10, 1, 3, 4, 7, 9, 12, 25, 55}, a)
}

func TestQuickSort(t *testing.T) {
	a := []int{25, 4, 12, -83, 7, 9, -10, 55, 3, 1}
	QuickSort(a, 0, len(a)-1)
	assert.Equal(t, []int{-83, -10, 1, 3, 4, 7, 9, 12, 25, 55}, a)
}
