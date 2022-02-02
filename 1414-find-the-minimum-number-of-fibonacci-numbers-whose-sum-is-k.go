package main

func findMinFibonacciNumbers(k int) int {
	fibonacci := []int{1, 1}
	for fibonacci[len(fibonacci)-1] < k {
		l := len(fibonacci) - 1
		fibonacci = append(fibonacci, fibonacci[l]+fibonacci[l-1])
	}
	var result int
	for i := len(fibonacci) - 1; k != 0; i-- {
		if k >= fibonacci[i] {
			k -= fibonacci[i]
			result++
		}
	}
	return result
}
