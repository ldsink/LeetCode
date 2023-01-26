package main

func greatestLetter(s string) string {
	counter := make([][2]int, 26)
	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			counter[r-'A'][0] = 1
		} else {
			counter[r-'a'][1] = 1
		}
	}
	for i := 25; i >= 0; i-- {
		if counter[i][0]+counter[i][1] == 2 {
			return string('A' + i)
		}
	}
	return ""
}
