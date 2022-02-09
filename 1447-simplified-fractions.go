package main

import "fmt"

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func simplifiedFractions(n int) []string {
	var result []string
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(j, i) == 1 {
				result = append(result, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return result
}
