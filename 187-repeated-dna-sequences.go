package main

func r2n(r uint8) int {
	if r == 'A' {
		return 0
	} else if r == 'C' {
		return 1
	} else if r == 'G' {
		return 2
	} else {
		return 3
	}
}

func findRepeatedDnaSequences(s string) []string {
	var result []string
	sequences := make(map[int]int)
	mask := 1<<20 - 1
	for i, j := 0, 0; i < len(s); i++ {
		j = ((j << 2) ^ r2n(s[i])) & mask
		if i >= 9 {
			sequences[j] += 1
			if sequences[j] == 2 {
				result = append(result, s[i-9:i+1])
			}
		}
	}
	return result
}
