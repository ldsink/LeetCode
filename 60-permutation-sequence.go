package main

func getPermutation(n int, k int) string {
	if n == 1 {
		return string(rune('0' + k))
	}
	count := 1
	for i := n - 1; i > 1; i-- {
		count *= i
	}
	offset := (k - 1) / count
	currentRune := rune('1' + offset)
	rest := getPermutation(n-1, ((k-1)%count)+1)
	restRunes := []rune(rest)
	for i := 0; i < len(restRunes); i++ {
		if restRunes[i] >= currentRune {
			restRunes[i] += 1
		}
	}
	restRunes = append([]rune{currentRune}, restRunes...)
	return string(restRunes)
}
