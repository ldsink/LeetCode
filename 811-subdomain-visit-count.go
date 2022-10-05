package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	visits := make(map[string]int)
	const DOT = "."
	const SPACE = " "
	for _, cpdomain := range cpdomains {
		parts := strings.Split(cpdomain, SPACE)
		count, _ := strconv.Atoi(parts[0])
		domain := parts[1]
		visits[domain] += count
		parts = strings.Split(domain, DOT)
		for i := 1; i < len(parts); i++ {
			visits[strings.Join(parts[i:], DOT)] += count
		}
	}
	var result []string
	for domain, count := range visits {
		result = append(result, fmt.Sprintf("%d %s", count, domain))
	}
	return result
}
