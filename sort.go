package main

func merge(leftArray, rightArray []int64) []int64 {
	var array []int64
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

func sort(array []int64) []int64 {
	if array[0] >= array[1] {
		return []int64{array[1], array[0]}
	}
	return array
}

func mergeSort(array []int64) []int64 {
	if len(array) > 2 {
		half := len(array) / 2
		return merge(mergeSort(array[:half]), mergeSort(array[half:]))
	} else if len(array) == 2 {
		return sort(array)
	}
	return array
}

func choosePivot(array []int64) int {
	if len(array) >= 3 {
		middle := len(array) / 2
		last := len(array) - 1
		if array[0] < array[middle] && array[middle] < array[last] {
			return middle
		} else if array[0] < array[last] && array[last] < array[middle] {
			return last
		}
	}
	return 0
}

func quickSort(array []int64) []int64 {
	if len(array) > 1 {
		pivot := choosePivot(array)

		i := 0
		for j := 0; j < len(array); j++ {
			if j != pivot && array[j] < array[pivot] {
				if i == pivot {
					pivot = j
				}
				swap := array[i]
				array[i] = array[j]
				array[j] = swap
				i++
			}
		}

		if pivot > i {
			swap := array[pivot]
			array[pivot] = array[i]
			array[i] = swap
		}

		leftPartion := quickSort(array[:i])
		rightPartition := quickSort(array[i+1:])
		return append(append(leftPartion, array[i]), rightPartition...)
	}
	return array
}
