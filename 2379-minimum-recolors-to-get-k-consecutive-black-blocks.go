package main

func minimumRecolors(blocks string, k int) int {
	result := len(blocks)
	for i := 0; i <= len(blocks)-k; i++ {
		var cost int
		for j := 0; j < k; j++ {
			if blocks[i+j] != 'B' {
				cost++
			}
		}
		if cost == 0 {
			return 0
		}
		if result > cost {
			result = cost
		}
	}
	return result
}
