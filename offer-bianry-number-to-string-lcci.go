package main

import (
	"fmt"
)

func printBin(num float64) string {
	var result []rune
	for i, j := 0, 0.5; i < 32 && num != 0; i, j = i+1, j/2 {
		if num >= j {
			num -= j
			result = append(result, '1')
		} else {
			result = append(result, '0')
		}
	}
	if len(result) == 32 {
		return "ERROR"
	}
	return fmt.Sprintf("0.%s", string(result))
}
