package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlgorithms(t *testing.T) {
	t.Run("quickSortAndBinarySearch", quickSortAndBinarySearch)
	t.Run("quickSortAndLinearSearch", quickSortAndLinearSearch)
	t.Run("mergeSortAndBinarySearch", mergeSortAndBinarySearch)
	t.Run("mergeSortAndLinearSearch", mergeSortAndLinearSearch)
	t.Run("bubbleSortAndBinarySearch", bubbleSortAndBinarySearch)
	t.Run("bubbleSortAndLinearSearch", bubbleSortAndLinearSearch)
	t.Run("fibonacciTest", fibonacciTest)
	t.Run("insertionSortAndBinarySearch", insertionSortAndBinarySearch)
	t.Run("insertionSortAndLinearSearch", insertionSortAndLinearSearch)
	t.Run("selectionSortAndBinarySearch", selectionSortAndBinarySearch)
	t.Run("selectionSortAndLinearSearch", selectionSortAndLinearSearch)
}

func createArrayForTest(numElements int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	randomSlice := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		var randomNum int
		for {
			randomNum = rand.Intn(numElements)
			if randomNum != 0 {
				randomSlice = append(randomSlice, randomNum)
				break
			}
		}
	}

	return randomSlice
}

func generateRandomIntForSearch(arr []int) int {
	var randomValue int
	for {
		randomValue = arr[rand.Intn(len(arr))]
		if randomValue != 0 {
			break
		}
	}
	return randomValue
}

func quickSortAndBinarySearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// quicksort
	quicksortStart := time.Now()
	t.Log("Starting quicksort")
	quickSort(arr, 0, len(arr)-1)
	quicksortEnd := time.Now()
	t.Log("Ending quicksort, with duration:", quicksortEnd.Sub(quicksortStart))

	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// binarySearch
	binarySearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting binarySearch for", value)
	found := binarySearch(arr, value, 0, len(arr)-1)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	binarySearchEnd := time.Now()
	t.Log("Ending binarySearch, with duration", binarySearchEnd.Sub(binarySearchStart))

	assert.Equal(t, arr, arrForCheck)
}

func quickSortAndLinearSearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// quicksort
	quicksortStart := time.Now()
	t.Log("Starting quicksort")
	quickSort(arr, 0, len(arr)-1)
	quicksortEnd := time.Now()
	t.Log("Ending quicksort, with duration:", quicksortEnd.Sub(quicksortStart))

	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// linearsearch
	linearSearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting linearsearch for", value)
	found := linearSearch(arr, value)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	linearSearchEnd := time.Now()
	t.Log("Ending linearsearch, with duration", linearSearchEnd.Sub(linearSearchStart))

	assert.Equal(t, arr, arrForCheck)
}

func mergeSortAndBinarySearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// mergeSort
	mergeSortStart := time.Now()
	t.Log("Starting mergeSort")
	sortedArray := mergeSort(arr)
	mergeSortEnd := time.Now()
	t.Log("Ending mergeSort, with duration:", mergeSortEnd.Sub(mergeSortStart))

	arrForCheck := make([]int, len(sortedArray))
	copy(arrForCheck, sortedArray)

	// binarySearch
	binarySearchStart := time.Now()
	value := generateRandomIntForSearch(sortedArray)
	t.Log("Starting binarySearch for", value)
	found := binarySearch(sortedArray, value, 0, len(sortedArray)-1)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	binarySearchEnd := time.Now()
	t.Log("Ending binarySearch, with duration", binarySearchEnd.Sub(binarySearchStart))

	assert.Equal(t, sortedArray, arrForCheck)
}

func mergeSortAndLinearSearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// mergeSort
	mergeSortStart := time.Now()
	t.Log("Starting mergeSort")
	sortedArray := mergeSort(arr)
	mergeSortEnd := time.Now()
	t.Log("Ending mergeSort, with duration:", mergeSortEnd.Sub(mergeSortStart))

	arrForCheck := make([]int, len(sortedArray))
	copy(arrForCheck, sortedArray)

	// linearsearch
	linearSearchStart := time.Now()
	value := generateRandomIntForSearch(sortedArray)
	t.Log("Starting linearsearch for", value)
	found := linearSearch(sortedArray, value)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	linearSearchEnd := time.Now()
	t.Log("Ending linearsearch, with duration", linearSearchEnd.Sub(linearSearchStart))

	assert.Equal(t, sortedArray, arrForCheck)
}

func bubbleSortAndBinarySearch(t *testing.T) {
	arr := createArrayForTest(10000)

	// bubbleSort
	bubbleSortStart := time.Now()
	t.Log("Starting bubbleSort")
	bubbleSort(arr)
	bubbleSortEnd := time.Now()
	t.Log("Ending bubbleSort, with duration:", bubbleSortEnd.Sub(bubbleSortStart))

	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// binarySearch
	binarySearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting binarySearch for", value)
	found := binarySearch(arr, value, 0, len(arr)-1)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	binarySearchEnd := time.Now()
	t.Log("Ending binarySearch, with duration", binarySearchEnd.Sub(binarySearchStart))

	assert.Equal(t, arr, arrForCheck)
}

func bubbleSortAndLinearSearch(t *testing.T) {
	arr := createArrayForTest(10000)

	// bubbleSort
	bubbleSortStart := time.Now()
	t.Log("Starting bubbleSort")
	bubbleSort(arr)
	bubbleSortEnd := time.Now()
	t.Log("Ending bubbleSort, with duration:", bubbleSortEnd.Sub(bubbleSortStart))

	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// linearsearch
	linearSearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting linearsearch for", value)
	found := linearSearch(arr, value)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	linearSearchEnd := time.Now()
	t.Log("Ending linearsearch, with duration", linearSearchEnd.Sub(linearSearchStart))

	assert.Equal(t, arr, arrForCheck)
}

func fibonacciTest(t *testing.T) {
	// fibonacci
	fibonacciStart := time.Now()
	t.Log("Starting fibonacci")
	fibonacciResult := fibonacci(30)
	fibonacciSimpleRes := fibonacciSimple(30)
	t.Log(fibonacciResult, fibonacciSimpleRes)
	fibonacciEnd := time.Now()
	t.Log("Ending fibonacci, with duration:", fibonacciEnd.Sub(fibonacciStart))
}

func insertionSortAndBinarySearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// insertionSort
	insertionSortStart := time.Now()
	t.Log("Starting insertionSort")
	insertionSort(arr, 0, 1)
	insertionSortEnd := time.Now()
	t.Log("Ending insertionSort, with duration:", insertionSortEnd.Sub(insertionSortStart))
	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// binarySearch
	binarySearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting binarySearch for", value)
	found := binarySearch(arr, value, 0, len(arr)-1)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	binarySearchEnd := time.Now()
	t.Log("Ending binarySearch, with duration", binarySearchEnd.Sub(binarySearchStart))

	assert.Equal(t, arr, arrForCheck)
}

func insertionSortAndLinearSearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// insertionSort
	insertionSortStart := time.Now()
	t.Log("Starting insertionSort")
	insertionSort(arr, 0, 1)
	insertionSortEnd := time.Now()
	t.Log("Ending insertionSort, with duration:", insertionSortEnd.Sub(insertionSortStart))
	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// linearsearch
	linearSearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting linearsearch for", value)
	found := linearSearch(arr, value)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	linearSearchEnd := time.Now()
	t.Log("Ending linearsearch, with duration", linearSearchEnd.Sub(linearSearchStart))

	assert.Equal(t, arr, arrForCheck)
}

func selectionSortAndBinarySearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// selectionSort
	selectionSortStart := time.Now()
	t.Log("Starting selectionSort")
	selectionSort(arr)
	selectionSortEnd := time.Now()
	t.Log("Ending selectionSort, with duration:", selectionSortEnd.Sub(selectionSortStart))

	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// binarySearch
	binarySearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting binarySearch for", value)
	found := binarySearch(arr, value, 0, len(arr)-1)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	binarySearchEnd := time.Now()
	t.Log("Ending binarySearch, with duration", binarySearchEnd.Sub(binarySearchStart))

	assert.Equal(t, arr, arrForCheck)
}

func selectionSortAndLinearSearch(t *testing.T) {
	arr := createArrayForTest(100000)

	// selectionSort
	selectionSortStart := time.Now()
	t.Log("Starting selectionSort")
	selectionSort(arr)
	selectionSortEnd := time.Now()
	t.Log("Ending selectionSort, with duration:", selectionSortEnd.Sub(selectionSortStart))

	arrForCheck := make([]int, len(arr))
	copy(arrForCheck, arr)

	// linearsearch
	linearSearchStart := time.Now()
	value := generateRandomIntForSearch(arr)
	t.Log("Starting linearsearch for", value)
	found := linearSearch(arr, value)
	if !found {
		t.Fatal("Value", value, "not found")
	}
	linearSearchEnd := time.Now()
	t.Log("Ending linearsearch, with duration", linearSearchEnd.Sub(linearSearchStart))

	assert.Equal(t, arr, arrForCheck)
}
