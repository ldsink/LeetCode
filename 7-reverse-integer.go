package main

import "fmt"

func reverse64(x int64) int64 {
	if x < 0 {
		r := -reverse64(-x)
		if r < -(1 << 31) {
			return 0
		}
		return r
	}
	var bitValues []int64
	for ; x != 0; {
		bitValues = append(bitValues, x%10)
		x = x / 10
	}
	var result int64
	for i, j := 1, 0; j < len(bitValues); i, j = 10*i, j+1 {
		result += int64(i) * bitValues[len(bitValues)-j-1]
	}
	if result > (1<<31)-1 {
		return 0
	}
	return result
}

func reverse(x int) int {
	return int(reverse64(int64(x)))
}

func main() {
	x := -2147483648
	fmt.Println(reverse(x))
}
