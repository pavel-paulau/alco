package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	array := []int64{1, -1, 7, 3, 2, 7, 6, 8, 5}
	assert.Equal(t, []int64{-1, 1, 2, 3, 5, 6, 7, 7, 8}, mergeSort(array))
}

func TestQuickSort(t *testing.T) {
	array := []int64{1, -1, 7, 3, 2, 7, 6, 8, 5}
	assert.Equal(t, []int64{-1, 1, 2, 3, 5, 6, 7, 7, 8}, quickSort(array))
}

func TestQuickSortFile(t *testing.T) {
	array := readArray("data/2sumSmall.txt")
	sortedArray := quickSort(array)

	for i, v := range sortedArray[:len(sortedArray)-2] {
		assert.True(t, v <= sortedArray[i+1], "Wrong order")
	}
}
