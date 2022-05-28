package main

import (
	"regexp"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	const neither = "Neither"
	const IPv4 = "IPv4"
	const IPv6 = "IPv6"
	const dot = "."
	const sep = ":"

	isIPv4 := func() bool {
		parts := strings.Split(queryIP, dot)
		if len(parts) != 4 { // 长度必须是 4
			return false
		}
		var address []string
		for _, part := range parts {
			if num, err := strconv.Atoi(part); err != nil || !(0 <= num && num <= 255) {
				return false
			} else {
				address = append(address, strconv.Itoa(num))
			}
		}
		return strings.Join(address, dot) == queryIP
	}

	re := regexp.MustCompile("^[\\da-fA-F]{1,4}$")
	isIPv6 := func() bool {
		parts := strings.Split(queryIP, sep)
		if len(parts) != 8 { // 长度必须是 8
			return false
		}
		for _, part := range parts {
			if !re.MatchString(part) {
				return false
			}
		}
		return true
	}

	if strings.Contains(queryIP, dot) && isIPv4() {
		return IPv4
	} else if strings.Contains(queryIP, sep) && isIPv6() {
		return IPv6
	}
	return neither
}
