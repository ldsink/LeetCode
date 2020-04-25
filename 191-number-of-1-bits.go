package main

func hammingWeight(num uint32) int {
	result := int(num & 0x1)
	for i := uint(0); i < 31; i++ {
		num >>= 1
		result += int(num & 0x1)
	}
	return result
}
