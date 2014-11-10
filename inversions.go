package main

func mergeAndCount(leftArray, rightArray []int) (int, []int) {
	var array []int
	var i, j, inversions int
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			array = append(array, leftArray[i])
			i++
		} else {
			array = append(array, rightArray[j])
			j++
			inversions += len(leftArray) - i
		}
	}
	if i == len(leftArray) {
		return inversions, append(array, rightArray[j:]...)
	}
	return inversions, append(array, leftArray[i:]...)
}

func sortAndCount(array []int) (int, []int) {
	if array[0] >= array[1] {
		return 1, []int{array[1], array[0]}
	}
	return 0, array
}

func countInversions(inversions int, array []int) (int, []int) {
	if len(array) > 2 {
		half := len(array) / 2
		leftInversions, leftArray := countInversions(inversions, array[:half])
		rightInversions, rightArray := countInversions(inversions, array[half:])
		splitInversions, sortedArray := mergeAndCount(leftArray, rightArray)
		return leftInversions + rightInversions + splitInversions, sortedArray
	} else if len(array) == 2 {
		return sortAndCount(array)
	} else {
		return 0, array
	}
}
