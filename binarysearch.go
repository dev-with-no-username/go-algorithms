package main

func binarySearch(sortedArray []int, valueToSearch int, leftIndex int, rightIndex int) bool {
	if rightIndex >= leftIndex {
		middle := (leftIndex + rightIndex) / 2
		if sortedArray[middle] == valueToSearch {
			return true
		}
		if sortedArray[middle] > valueToSearch {
			binarySearch(sortedArray, valueToSearch, leftIndex, middle-1)
		}
		if sortedArray[middle] < valueToSearch {
			binarySearch(sortedArray, valueToSearch, middle+1, rightIndex)
		}
	} else {
		return false
	}
	// Condition base that I set to true to not cause unwanted behaviour
	// in case of this will be returned before the actual "return"
	// encountered by the above "if" statement of the algorithm
	return true
}
