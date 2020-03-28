package main

func isZero(n, b int) bool {
	return (n>>uint(b))&0x1 == 0x0
}

func xorBit(n, b int) int {
	return n ^ (0x1 << uint(b))
}

func grayCode(n int) []int {
	result := []int{0}
	if n == 0 {
		return result
	}
	used := make(map[int]bool)
	for changed, num := true, 0; changed; {
		changed = false
		for i := 0; i < n; i++ {
			if isZero(num, i) {
				if r := xorBit(num, i); !used[r] {
					num = r
					used[num] = true
					result = append(result, num)
					changed = true
					break
				}
			} else {
				if r := xorBit(num, i); r > 0 && !used[r] {
					num = r
					used[num] = true
					result = append(result, num)
					changed = true
					break
				}
			}
		}
	}
	return result
}
