package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterativeMultiplication(t *testing.T) {
	a := [][]int{
		[]int{1, 2, 5, 4},
		[]int{1, 1, 3, 8},
		[]int{7, 3, 6, 9},
		[]int{2, 5, 4, 7},
	}
	b := [][]int{
		[]int{9, 8, 1, 7},
		[]int{2, 7, 2, 8},
		[]int{1, 5, 4, 3},
		[]int{2, 3, 6, 7},
	}
	c := [][]int{
		[]int{26, 59, 49, 66},
		[]int{30, 54, 63, 80},
		[]int{93, 134, 91, 154},
		[]int{46, 92, 70, 115},
	}
	assert.Equal(t, c, multiplyIterative(a, b))
}
