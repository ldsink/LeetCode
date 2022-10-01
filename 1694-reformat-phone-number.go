package main

import (
	"strings"
)

func reformatNumber(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	if len(number) < 3 {
		return number
	}

	var result []string
	if offset := len(number) % 3; offset == 0 {
		for i := 0; i < len(number); i += 3 {
			result = append(result, number[i:i+3])
		}
	} else if offset == 1 {
		for i := 0; i < len(number)-4; i += 3 {
			result = append(result, number[i:i+3])
		}
		result = append(result, number[len(number)-4:len(number)-2])
		result = append(result, number[len(number)-2:])
	} else {
		for i := 0; i < len(number)-2; i += 3 {
			result = append(result, number[i:i+3])
		}
		result = append(result, number[len(number)-2:])
	}
	return strings.Join(result, "-")
}
