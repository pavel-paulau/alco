package alco

func mergeAndCount(leftArray, rightArray []int) (int, []int) {
	var array []int
	var i, j, inversions int
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			array = append(array, leftArray[i])
			i += 1
		} else {
			array = append(array, rightArray[j])
			j += 1
			inversions += len(leftArray) - i
		}
	}
	if i == len(leftArray) {
		return inversions, append(array, rightArray[j:]...)
	} else {
		return inversions, append(array, leftArray[i:]...)
	}
}

func sortAndCount(array []int) (int, []int) {
	if array[0] >= array[1] {
		return 1, []int{array[1], array[0]}
	} else {
		return 0, array
	}
}

func CountInversions(inversions int, array []int) (int, []int) {
	if len(array) > 2 {
		half := len(array) / 2
		leftInversions, leftArray := CountInversions(inversions, array[:half])
		rightInversions, rightArray := CountInversions(inversions, array[half:])
		splitInversions, sortedArray := mergeAndCount(leftArray, rightArray)
		return leftInversions + rightInversions + splitInversions, sortedArray
	} else if len(array) == 2 {
		return sortAndCount(array)
	} else {
		return 0, array
	}
}
