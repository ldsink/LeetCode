package main

func decrypt(code []int, k int) []int {
	if k == 0 {
		for i := 0; i < len(code); i++ {
			code[i] = 0
		}
		return code
	}
	var getNewCode func(int) int
	if k > 0 {
		getNewCode = func(i int) (r int) {
			for j := 1; j <= k; j++ {
				idx := (i + j) % len(code)
				r += code[idx]
			}
			return
		}
	} else {
		getNewCode = func(i int) (r int) {
			for j := -1; j >= k; j-- {
				idx := (i + j) % len(code)
				if idx < 0 {
					idx += len(code)
				}
				r += code[idx]
			}
			return
		}
	}

	result := make([]int, len(code))
	for i := 0; i < len(code); i++ {
		result[i] = getNewCode(i)
	}
	for i := 0; i < len(code); i++ {
		code[i] = result[i]
	}
	return code
}
