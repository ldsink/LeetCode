package main

func calculateTax(brackets [][]int, income int) float64 {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var tax float64
	var prev int
	for _, bracket := range brackets {
		upper, percent := bracket[0], bracket[1]
		tax += float64((min(upper, income)-prev)*percent) / 100
		if upper >= income {
			break
		} else {
			prev = upper
		}
	}
	return tax
}
