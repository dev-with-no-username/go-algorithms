package main

func linearSearch(sortedArray []int, valueToSearch int) bool {
	for i := 0; i < len(sortedArray)-1; i++ {
		if sortedArray[i] == valueToSearch {
			return true
		}
	}
	return false
}
