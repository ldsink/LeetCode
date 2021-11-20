package main

import (
	"strings"
	"unicode"
)

func licenseKeyFormatting(s string, k int) string {
	var rs []rune
	const dash = '-'
	for _, r := range []rune(s) {
		if r == dash {
			continue
		}
		rs = append(rs, unicode.ToUpper(r))
	}
	var parts []string
	start := len(rs) % k
	if start != 0 {
		parts = append(parts, string(rs[:start]))
	}
	for ; start < len(rs); start += k {
		parts = append(parts, string(rs[start:start+k]))
	}
	return strings.Join(parts, string(dash))
}
