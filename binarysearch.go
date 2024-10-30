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
	return true
}
