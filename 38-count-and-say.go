package main

func countAndSay(n int) string {
	rs := []rune{'1'}
	for i := 1; i < n; i++ {
		var current []rune
		var count int
		var lastNum rune
		for _, r := range rs {
			if r == lastNum {
				count += 1
				continue
			}
			if count > 0 {
				current = append(current, rune('0'+count))
				current = append(current, lastNum)
			}
			lastNum = r
			count = 1
		}
		if count > 0 {
			current = append(current, rune('0'+count))
			current = append(current, lastNum)
		}
		rs = current
	}
	return string(rs)
}
