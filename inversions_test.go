package alco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountInversionsReversedArray(t *testing.T) {
	array := []int{6, 5, 4, 3, 2, 1}
	inversions, _ := CountInversions(0, array)
	assert.Equal(t, 15, inversions)
}

func TestCountInversionsAllSplit(t *testing.T) {
	array := []int{1, 3, 5, 2, 4, 6}
	inversions, _ := CountInversions(0, array)
	assert.Equal(t, 3, inversions)
}

func TestCountInversionsMixed(t *testing.T) {
	array := []int{1, 7, 6, 3, 3, 5}
	inversions, _ := CountInversions(0, array)
	assert.Equal(t, 7, inversions)
}
