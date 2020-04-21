package main

func r2n(r rune) int {
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
	sequences := make(map[int]int)
	funnel := 1<<20 - 1
	for i, j := 0, 0; i < len(s); i++ {
		j = ((j << 2) ^ r2n(rune(s[i]))) & funnel
		if i >= 9 {
			sequences[j] += 1
		}
	}
	nucleotides := []rune("ACGT")
	var result []string
	for sequence, count := range sequences {
		if count > 1 {
			var rs []rune
			for i := 18; i >= 0; i -= 2 {
				rs = append(rs, nucleotides[(sequence>>uint(i))&0x3])
			}
			result = append(result, string(rs))
		}
	}
	return result
}
