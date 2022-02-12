package main

func maxNumberOfBalloons(text string) int {
	count := make([]int, 26)
	for _, r := range []rune(text) {
		count[r-'a']++
	}
	// balloon
	max := count['b'-'a']
	if max > count[0] {
		max = count[0]
	}
	if max > count['l'-'a']/2 {
		max = count['l'-'a'] / 2
	}
	if max > count['o'-'a']/2 {
		max = count['o'-'a'] / 2
	}
	if max > count['n'-'a'] {
		max = count['n'-'a']
	}
	return max
}
