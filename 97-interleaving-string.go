package main

func dp(s1, s2, s3 string, saved map[int]map[int]bool) bool {
	if _, ok := saved[len(s1)]; !ok {
		saved[len(s1)] = make(map[int]bool)
	}
	if r, ok := saved[len(s1)][len(s2)]; ok {
		return r
	}

	var result bool
	if len(s1) == 0 {
		result = s2 == s3
	} else if len(s2) == 0 {
		result = s1 == s3
	} else {
		result = (s1[0] == s3[0] && dp(s1[1:], s2, s3[1:], saved)) || (
			s2[0] == s3[0] && dp(s1, s2[1:], s3[1:], saved))
	}
	saved[len(s1)][len(s2)] = result
	return result
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	saved := make(map[int]map[int]bool)
	return dp(s1, s2, s3, saved)
}
