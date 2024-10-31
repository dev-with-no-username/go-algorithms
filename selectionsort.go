package main

func selectionSort(arrToSort []int) {
	// set "i < len(arrToSort)-1", because j starts at i+1 due to the
	// fact that elements before j are already sorted
	for i := 0; i < len(arrToSort)-1; i++ {

		// set index of miminum value = i
		minIndex := i

		// loop through the entire array checking if there is some values
		// lower than the minimum and if yes changing the index of minimum
		for j := i + 1; j < len(arrToSort); j++ {
			if arrToSort[j] < arrToSort[minIndex] {
				minIndex = j
			}
		}

		// in case of values lower that the minimum, minIndex is changed
		// so we have to swap these values in-place so that at "i" index
		// there will be the arrToSort lowest value ever and in the next
		// inner for loop, we can start from the "i+1" index ignoring
		// previous values
		if minIndex != i {
			selectionSwap(arrToSort, minIndex, i)
		}
	}
}

func selectionSwap(arr []int, currentIndex int, destinationIndex int) {
	temp := arr[currentIndex]
	arr[currentIndex] = arr[destinationIndex]
	arr[destinationIndex] = temp
}
