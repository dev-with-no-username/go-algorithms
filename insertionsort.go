package main

func insertionSort(arrToSort []int, firstIndex int, secondIndex int) {
	// stop when all elements have already been compared and we are past the
	// end of the array, cause the len() returns the elements number + 1
	if secondIndex == len(arrToSort) {
		return
	}

	// initialize two indexes to be able to compare arrToSort[secondIndex] with
	// each values before it, changing in place without moving starting indexes
	tempFirstIndex := firstIndex
	tempSecondIndex := secondIndex
	for tempFirstIndex >= 0 && tempSecondIndex < len(arrToSort) {
		if arrToSort[tempSecondIndex] < arrToSort[tempFirstIndex] {
			insertionSwap(arrToSort, tempFirstIndex, tempSecondIndex)
			tempSecondIndex = tempFirstIndex
		}
		tempFirstIndex -= 1
	}

	// start recursion moving to right (+1) the secondIndex and starting the
	// comparison from current secondIndex that will be the first previous
	// element of next iteration
	insertionSort(arrToSort, secondIndex, secondIndex+1)
}

func insertionSwap(arr []int, srcIndex int, dstIndex int) {
	temp := arr[srcIndex]
	arr[srcIndex] = arr[dstIndex]
	arr[dstIndex] = temp
}
