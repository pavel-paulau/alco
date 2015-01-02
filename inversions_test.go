package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountInversionsReversedArray(t *testing.T) {
	array := []int64{6, 5, 4, 3, 2, 1}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 15, inversions)
}

func TestCountInversionsAllSplit(t *testing.T) {
	array := []int64{1, 3, 5, 2, 4, 6}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 3, inversions)
}

func TestCountInversionsMixed(t *testing.T) {
	array := []int64{1, 7, 6, 3, 3, 5}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 7, inversions)
}

/*
This file contains all of the 100,000 integers between 1 and 100,000 (inclusive)
in some order, with no integer repeated.

Your task is to compute the number of inversions in the file given, where the
ith row of the file indicates the ith entry of an array.
*/
func TestCountInversions(t *testing.T) {
	array := readArray("data/IntegerArray.txt")
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 2407905288, inversions)
}
