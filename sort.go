package alco

func merge(array1, array2 []int) []int {
	var array []int
	var i, j int
	for i < len(array1) && j < len(array2) {
		if array1[i] < array2[j] {
			array = append(array, array1[i])
			i += 1
		} else {
			array = append(array, array2[j])
			j += 1
		}
	}
	if i == len(array1) {
		return append(array, array2[j:]...)
	} else {
		return append(array, array1[i:]...)
	}
}

func sort2(array []int) []int {
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
		return sort2(array)
	} else {
		return array
	}
}
