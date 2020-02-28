package main

import (
	"fmt"
	"unsafe"
)

func getBit(num, n uint) uint {
	return num >> n & 1
}

func uintIsTrue(i uint) bool {
	return !(i == 0)
}

func getUIntSum(a, b uint) uint {
	var result, carry, bit uint
	for i, j := uint(0), uint(unsafe.Sizeof(result)<<3); i < j; i++ {
		aBit := getBit(a, i)
		bBit := getBit(b, i)
		if aBit != bBit {
			if uintIsTrue(carry) {
				bit = 0
			} else {
				bit = 1
			}
		} else if aBit == 1 {
			if uintIsTrue(carry) {
				bit = 1
			} else {
				bit = 0
				carry = 1
			}
		} else {
			if uintIsTrue(carry) {
				bit = 1
				carry = 0
			} else {
				bit = 0
			}
		}
		result = result ^ (bit << i)
	}
	return result
}

func getSum(a int, b int) int {
	return int(getUIntSum(uint(a), uint(b)))
}

func main() {
	result := getSum(1, 2)
	fmt.Println(result)

	result = getSum(-1, 3)
	fmt.Println(result)

	result = getSum(-123, 0)
	fmt.Println(result)

	result = getSum(-123, -4)
	fmt.Println(result)
}
