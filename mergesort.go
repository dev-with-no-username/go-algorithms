package main

func mergeSort(items []int) []int {
	// Base case. A list of zero or one elements is sorted, by definition.
	if len(items) < 2 {
		return items
	}

	// Recursive case. Divide the list into equal-sized sublists consisting
	// of the first half and second half of the list and sort both
	firstHalf := mergeSort(items[:len(items)/2])
	secondHalf := mergeSort(items[len(items)/2:])

	// Merge the sorted sublists
	return merge(firstHalf, secondHalf)
}

func merge(firstHalf []int, secondHalf []int) []int {
	sortedArray := []int{}
	i := 0
	j := 0

	// Loop through each element of both halves, comparing each element of one
	// half with each element of the other half, placing them in a new list
	// in ascending order
	for i < len(firstHalf) && j < len(secondHalf) {
		if firstHalf[i] < secondHalf[j] {
			sortedArray = append(sortedArray, firstHalf[i])
			i++
		} else {
			sortedArray = append(sortedArray, secondHalf[j])
			j++
		}
	}

	// In case of halves are made by different number of elements, adding
	// remaining elements that are sorted for sure: in this implementation
	// it could be one single element that will be the maximum value
	for ; i < len(firstHalf); i++ {
		sortedArray = append(sortedArray, firstHalf[i])
	}
	for ; j < len(secondHalf); j++ {
		sortedArray = append(sortedArray, secondHalf[j])
	}

	return sortedArray
}
