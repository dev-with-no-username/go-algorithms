package main

func fibonacciSimple(num int) int {
	oldSum := 0
	sum := 0
	for i := 0; i <= num; i++ {
		if i == 0 {
			sum += 0
		} else if i == 1 {
			sum += 1
		} else {
			temp := sum
			sum += oldSum
			oldSum = temp
		}
	}
	return sum
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}
