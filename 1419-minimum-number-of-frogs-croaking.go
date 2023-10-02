package main

func minNumberOfFrogs(croakOfFrogs string) int {
	step := make(map[rune]int)
	for i, r := range []rune("croak") {
		step[r] = i
	}
	var count [5]int // count[4] 代表可用的青蛙
	var frogs int
	for _, r := range []rune(croakOfFrogs) {
		i := step[r]
		if i == 0 {
			if count[4] == 0 {
				frogs++
				count[4]++
			}
			count[i]++
			count[4]--
		} else if count[i-1] == 0 {
			return -1
		} else {
			count[i-1]--
			count[i]++
		}
	}
	for i := 0; i < 4; i++ {
		if count[i] != 0 {
			return -1
		}
	}
	return frogs
}
