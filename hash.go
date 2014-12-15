package main

func find2SumBF(array []int64, sum int64) int64 {
	counter := int64(0)
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == sum {
				counter++
			}
		}
	}
	return counter
}

func find2Sum(array []int64, sum int64) int64 {
	sumHash := map[int64]int64{}
	for i := 0; i < len(array); i++ {
		if _, ok := sumHash[array[i]]; ok {
			sumHash[array[i]]++
		} else {
			sumHash[array[i]] = 1
		}
	}
	counter := int64(0)
	for number, firstCounter := range sumHash {
		if number <= sum/2 {
			if secondCounter, ok := sumHash[sum-number]; ok {
				if number == sum-number {
					counter += firstCounter * (firstCounter - 1) / 2
				} else {
					counter += firstCounter * secondCounter
				}
			}
		}
	}
	return counter
}

func find2SumAny(array []int64, sum int64) int64 {
	sumHash := map[int64]int64{}
	for i := 0; i < len(array); i++ {
		sumHash[array[i]] = 1
	}
	counter := int64(0)
	for number := range sumHash {
		if number <= sum/2 {
			if _, ok := sumHash[sum-number]; ok {
				counter++
				break
			}
		}
	}
	return counter
}

func findWindow(array []int64, window int64, initIdx int) int {
	half := len(array) / 2
	if half == 0 || window == array[half] {
		return half + initIdx
	}

	if window > array[half] {
		return findWindow(array[half:], window, half+initIdx)
	}
	return findWindow(array[:half], window, initIdx)
}

func findExistingSums(array []int64, tMin, tMax int64) int {
	array = mergeSort(array)
	min := array[0]
	max := array[len(array)-1]
	sums := map[int64]bool{}
	var window int
	for _, v := range array {
		if min <= tMin-v && tMin-v <= max {
			window = findWindow(array, tMin-v, 0)
			for window < len(array) && v+array[window] <= tMax {
				if v+array[window] >= tMin && v != array[window] {
					sums[v+array[window]] = true
				}
				window++
			}
		}
	}
	return len(sums)
}
