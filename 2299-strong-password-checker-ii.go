package main

import (
	"strings"
	"unicode"
)

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	var hasLower, hasUpper, hasDigit, hasSpecial bool
	for i, ch := range password {
		if i+1 < len(password) && password[i] == password[i+1] {
			return false
		}
		if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else if strings.ContainsRune("!@#$%^&*()-+", ch) {
			hasSpecial = true
		}
	}
	return hasLower && hasUpper && hasDigit && hasSpecial
}
