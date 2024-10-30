package main

func bubbleSort(arrToSort []int) {
	changing := true
	for changing {
		changing = false
		for i := 0; i < len(arrToSort)-1; i++ {
			if arrToSort[i] > arrToSort[i+1] {
				temp := arrToSort[i]
				arrToSort[i] = arrToSort[i+1]
				arrToSort[i+1] = temp
				changing = true
			}
		}
	}
}
