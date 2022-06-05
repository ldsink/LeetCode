package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	count := make(map[string]int)
	for _, email := range emails {
		parts := strings.SplitN(email, "@", 2)
		localName, domain := parts[0], parts[1]
		if strings.Contains(localName, "+") {
			parts := strings.SplitN(localName, "+", 2)
			localName = parts[0]
		}
		localName = strings.Replace(localName, ".", "", -1)
		email = fmt.Sprintf("%s@%s", localName, domain)
		count[email]++
	}
	return len(count)
}
