package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2SumSmall(t *testing.T) {
	array := []int64{1, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 5, 6, 8}
	pairs := find2Sum(array, 6)
	assert.Equal(t, find2SumBF(array, 6), pairs)

	pairs = find2Sum(array, 7)
	assert.Equal(t, find2SumBF(array, 7), pairs)
}

func readArray(fname string) []int64 {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	array := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		element, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		array = append(array, element)
	}
	return array
}

func Test2Sum(t *testing.T) {
	array := readArray("data/2sum.txt")
	pairs := find2Sum(array, 0)
	assert.Equal(t, 0, pairs)
}

func Test2SumRangeSmall(t *testing.T) {
	array := []int64{-28, -15, -13, -13, -9, 8, 9, 10, 12, 12, 18}
	sums := int64(0)
	for sum := -2; sum <= 2; sum++ {
		sums += find2SumAny(array, int64(sum))
	}
	// -13 + 12 = -1 or -9 + 8 = -1
	// -9 + 9 = 0
	// -9 + 10 = 1
	assert.Equal(t, 3, sums)
}

/*
The file contains 1 million integers, both positive and negative (there might be
some repetitions).

The task is to compute the number of target values t in the interval
[-10000,10000] (inclusive) such that there are distinct numbers x,y in the input
file that satisfy x+y=t.
*/
func TestFindExistingSumsSmall(t *testing.T) {
	array := []int64{-28, -15, -13, -13, -9, 8, 9, 10, 12, 12, 18}
	sums := findExistingSums(array, -2, 2)
	assert.Equal(t, 3, sums)

	array = []int64{-10001, 1, 2, -10001}
	sums = findExistingSums(array, -10000, 10000)
	assert.Equal(t, 3, sums)

	array = []int64{-10001, 1, 2, -10001, 9999}
	sums = findExistingSums(array, -10000, 10000)
	assert.Equal(t, 5, sums)
}

func TestFindExistingSumsSmallFile(t *testing.T) {
	array := readArray("data/2sumSmall.txt")
	sums := findExistingSums(array, -10000, 10000)
	assert.Equal(t, 6, sums)
}

func TestFindExistingSums(t *testing.T) {
	array := readArray("data/2sum.txt")
	sums := findExistingSums(array, -10000, 10000)
	assert.Equal(t, 427, sums)
}
