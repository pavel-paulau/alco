package alco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	array := []int{1, -1, 7, 3, 2, 7, 6, 8, 5}
	assert.Equal(t, []int{-1, 1, 2, 3, 5, 6, 7, 7, 8}, MergeSort(array))
}
