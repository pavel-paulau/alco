package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
The goal of this problem is to implement the "Median Maintenance" algorithm.

The text file contains a list of the integers from 1 to 10000 in unsorted order.

Letting xi denote the ith number of the file, the kth median mk is defined as
the median of the numbers x1,…,xk. (So, if k is odd, then mk is ((k+1)/2)th
smallest number among x1,…,xk; if k is even, then mk is the (k/2)th smallest
number among x1,…,xk.)
*/

func TestMedianMaintenanceSmall(t *testing.T) {
	array := []int64{4, 3, 2, 9, 1}
	medians := medianMaintenance(array)
	assert.Equal(t, 16, medians)
}

func TestMedianMaintenance(t *testing.T) {
	array := readArray("data/Median.txt")
	medians := medianMaintenance(array)
	assert.Equal(t, 1213, medians%10000)
}
