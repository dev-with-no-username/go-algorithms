package main

func mergeSort(items []int) []int {
	// exit condition
	if len(items) < 2 {
		return items
	}

	firstHalf := mergeSort(items[:len(items)/2])
	secondHalf := mergeSort(items[len(items)/2:])
	return merge(firstHalf, secondHalf)
}

func merge(firstHalf []int, secondHalf []int) []int {
	sortedArray := []int{}
	i := 0
	j := 0
	for i < len(firstHalf) && j < len(secondHalf) {
		if firstHalf[i] < secondHalf[j] {
			sortedArray = append(sortedArray, firstHalf[i])
			i++
		} else {
			sortedArray = append(sortedArray, secondHalf[j])
			j++
		}
	}

	// in case of halves made by different number of elements
	for ; i < len(firstHalf); i++ {
		sortedArray = append(sortedArray, firstHalf[i])
	}
	for ; j < len(secondHalf); j++ {
		sortedArray = append(sortedArray, secondHalf[j])
	}

	return sortedArray
}
