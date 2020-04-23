package main

func reverseBits(num uint32) uint32 {
	var result uint32
	for i, j := uint(0), uint(31); i < 32; i, j = i+1, j-1 {
		result |= ((num >> j) & 0x1) << i
	}
	return result
}
