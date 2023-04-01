package main

import (
	"fmt"
	"strings"
	"unicode"
)

func maskPII(s string) string {
	if strings.Contains(s, "@") {
		parts := strings.Split(s, "@")
		username := strings.ToLower(parts[0])
		domain := strings.ToLower(parts[1])
		return fmt.Sprintf("%s*****%s@%s", username[:1], username[len(username)-1:], domain)
	}
	var numbers []rune
	for _, r := range s {
		if unicode.IsNumber(r) {
			numbers = append(numbers, r)
		}
	}
	num := string(numbers)
	phone := fmt.Sprintf("***-***-%s", num[len(num)-4:])
	if len(num) > 10 {
		var countryCode string
		for i := 10; i < len(num); i++ {
			countryCode += "*"
		}
		phone = fmt.Sprintf("+%s-%s", countryCode, phone)
	}
	return phone
}
