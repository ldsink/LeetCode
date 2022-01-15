package main

func totalMoney(n int) int {
	var total, base int
	for n > 0 {
		if n >= 7 {
			n -= 7
			total += 7*base + 28
			base++
		} else {
			for ; n > 0; n-- {
				total += base + n
			}
		}
	}
	return total
}
