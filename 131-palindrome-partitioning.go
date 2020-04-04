package main

func _partition(s string, si int, isPalindrome map[int]map[int]bool) [][]string {
	var result [][]string
	if len(s[si:]) == 0 {
		return result
	} else if len(s[si:]) == 1 {
		return append(result, []string{s[si:]})
	}

	for i := si + 1; i <= len(s); i++ {
		// ensure substring is palindrome
		if _, ok := isPalindrome[si]; !ok {
			isPalindrome[si] = make(map[int]bool)
		}
		if _, ok := isPalindrome[si][i]; !ok {
			isPalindrome[si][i] = true
			for j, k := si, i-1; j < k; j, k = j+1, k-1 {
				if s[j] != s[k] {
					isPalindrome[si][i] = false
					break
				}
			}
		}
		if !isPalindrome[si][i] {
			continue
		}

		if i == len(s) {
			result = append(result, []string{s[si:]})
			continue
		}

		r := _partition(s, i, isPalindrome)
		if len(r) == 0 {
			continue
		}
		current := []string{s[si:i]}
		for _, ss := range r {
			result = append(result, append(current, ss...))
		}
	}

	return result
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	isPalindrome := make(map[int]map[int]bool)
	return _partition(s, 0, isPalindrome)
}
