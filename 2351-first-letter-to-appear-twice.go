package main

func repeatedCharacter(s string) byte {
	exist := make([]bool, 26)
	for _, c := range s {
		if exist[c-'a'] {
			return byte(c)
		} else {
			exist[c-'a'] = true
		}
	}
	return 0
}
