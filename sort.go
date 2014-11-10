package main

func merge(leftArray, rightArray []int) []int {
	var array []int
	var i, j int
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			array = append(array, leftArray[i])
			i++
		} else {
			array = append(array, rightArray[j])
			j++
		}
	}
	if i == len(leftArray) {
		return append(array, rightArray[j:]...)
	}
	return append(array, leftArray[i:]...)
}

func sort(array []int) []int {
	if array[0] >= array[1] {
		return []int{array[1], array[0]}
	}
	return array
}

func mergeSort(array []int) []int {
	if len(array) > 2 {
		half := len(array) / 2
		return merge(mergeSort(array[:half]), mergeSort(array[half:]))
	} else if len(array) == 2 {
		return sort(array)
	} else {
		return array
	}
}

func choosePivot(array []int) int {
	return array[0]
}

func quickSort(array []int) []int {
	if len(array) > 1 {
		pivot := choosePivot(array)

		i := 1
		for j := 1; j < len(array); j++ {
			if array[j] < pivot {
				swap := array[i]
				array[i] = array[j]
				array[j] = swap
				i++
			}
		}

		swap := array[0]
		array[0] = array[i-1]
		array[i-1] = swap

		return append(
			append(quickSort(array[:i-1]), pivot),
			quickSort(array[i:])...)
	} else {
		return array
	}
}
