package main

import "math/rand"

func quickSort(arrToSort []int, lowIndex int, highIndex int) {
	if lowIndex >= highIndex {
		return
	}

	// + 1 because Intn return a random number inside [0,n) interval and the highIndex is len()-1
	pivot := arrToSort[rand.Intn(highIndex+1)]
	leftPointer := lowIndex
	rightPointer := highIndex

	for leftPointer < rightPointer {
		// Walk from the left until we find a number greater than the pivot, or hit the right pointer.
		for arrToSort[leftPointer] <= pivot && leftPointer < rightPointer {
			leftPointer += 1
		}
		// Walk from the right until we find a number less than the pivot, or hit the left pointer.
		for arrToSort[rightPointer] >= pivot && leftPointer < rightPointer {
			rightPointer -= 1
		}
		swap(arrToSort, leftPointer, rightPointer)
	}

	swap(arrToSort, leftPointer, highIndex)

	quickSort(arrToSort, lowIndex, leftPointer-1)
	quickSort(arrToSort, leftPointer+1, highIndex)
}

func swap(arr []int, currentIndex int, destinationIndex int) {
	temp := arr[currentIndex]
	arr[currentIndex] = arr[destinationIndex]
	arr[destinationIndex] = temp
}
