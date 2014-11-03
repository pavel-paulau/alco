package alco

func merge(leftArray, rightArray []int) []int {
	var array []int
	var i, j int
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			array = append(array, leftArray[i])
			i += 1
		} else {
			array = append(array, rightArray[j])
			j += 1
		}
	}
	if i == len(leftArray) {
		return append(array, rightArray[j:]...)
	} else {
		return append(array, leftArray[i:]...)
	}
}

func sort(array []int) []int {
	if array[0] >= array[1] {
		return []int{array[1], array[0]}
	} else {
		return array
	}
}

func MergeSort(array []int) []int {
	if len(array) > 2 {
		half := len(array) / 2
		return merge(MergeSort(array[:half]), MergeSort(array[half:]))
	} else if len(array) == 2 {
		return sort(array)
	} else {
		return array
	}
}
