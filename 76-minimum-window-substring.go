package main

func minWindow(s string, t string) string {
	tr := []rune(t)
	characters := make(map[rune]int)
	for _, r := range tr {
		characters[r] += 1
	}

	position := make(map[rune][]int)
	sr := []rune(s)
	for i, r := range sr {
		if characters[r] == 0 {
			continue
		}
		position[r] = append(position[r], i)
	}

	if len(position) < len(characters) {
		return ""
	}

	left, right := len(sr)-1, 0
	for r, v := range position {
		if len(v) < characters[r] {
			return ""
		}
		if left > v[0] {
			left = v[0]
		}
		if right < v[characters[r]-1] {
			right = v[characters[r]-1]
		}
	}

	for l, r := left, right; ; {
		leftRune := sr[l]
		if len(position[leftRune]) == characters[leftRune] {
			break
		}

		var newL int
		for newL = l + 1; newL < len(sr); newL++ {
			if characters[sr[newL]] > 0 {
				break
			}
		}
		position[leftRune] = position[leftRune][1:]
		if position[leftRune][characters[leftRune]-1] > r {
			r = position[leftRune][characters[leftRune]-1]
		}
		l = newL
		if (r - l) < (right - left) {
			left, right = l, r
		}
	}
	return s[left : right+1]
}
