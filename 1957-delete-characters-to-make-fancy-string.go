package main

func makeFancyString(s string) string {
	var result []rune
	rs := []rune(s)
	var r rune
	var count int
	for i := 0; i < len(rs); i++ {
		if count == 0 {
			count++
			r = rs[i]
		} else if rs[i] == r {
			if count < 2 {
				count++
			}
		} else {
			for ; count > 0; count-- {
				result = append(result, r)
			}
			count = 1
			r = rs[i]
		}
	}
	for ; count > 0; count-- {
		result = append(result, r)
	}
	return string(result)
}
