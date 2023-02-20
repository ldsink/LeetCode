package main

func bestHand(ranks []int, suits []byte) string {
	// Flush
	flush := true
	for i := 1; i < 5; i++ {
		if suits[i] != suits[0] {
			flush = false
			break
		}
	}
	if flush {
		return "Flush"
	}
	// Three of a Kind
	counter := make(map[int]int)
	for _, rank := range ranks {
		counter[rank]++
	}
	var maxCount int
	for _, v := range counter {
		if maxCount < v {
			maxCount = v
		}
	}
	if maxCount >= 3 {
		return "Three of a Kind"
	} else if maxCount == 2 {
		return "Pair"
	}
	return "High Card"
}
