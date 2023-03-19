package main

func findLexSmallestString(s string, a int, b int) string {
	result := s
	for i := 0; i < len(s); i++ {
		s = s[len(s)-b:] + s[:len(s)-b]
		for j, rs := 0, []rune(s); j < 10; j++ {
			for k := 1; k < len(s); k += 2 {
				rs[k] += rune(a)
				if rs[k] > '9' {
					rs[k] -= 10
				}
			}
			if b&1 == 1 {
				for p := 0; p < 10; p++ {
					for k := 0; k < len(s); k += 2 {
						rs[k] += rune(a)
						if rs[k] > '9' {
							rs[k] -= 10
						}
					}
					if result > string(rs) {
						result = string(rs)
					}
				}
			} else {
				if result > string(rs) {
					result = string(rs)
				}
			}
		}
	}
	return result
}
