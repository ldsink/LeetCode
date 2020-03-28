package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var result []string
	for i := 1; i <= 3 && i < len(s); i++ {
		if n, err := strconv.ParseInt(s[:i], 10, 64); err != nil || n > 255 {
			continue
		}
		if len(s[:i]) > 1 && s[0] == '0' {
			continue
		}
		for j, jp := 0, i+1; j < 3 && jp < len(s); j, jp = j+1, jp+1 {
			if n, err := strconv.ParseInt(s[i:jp], 10, 64); err != nil || n > 255 {
				continue
			}
			if len(s[i:jp]) > 1 && s[i] == '0' {
				continue
			}
			for k, kp := 0, jp+1; k < 3 && kp < len(s); k, kp = k+1, kp+1 {
				if n, err := strconv.ParseInt(s[jp:kp], 10, 64); err != nil || n > 255 {
					continue
				}
				if len(s[jp:kp]) > 1 && s[jp] == '0' {
					continue
				}
				if l := len(s[kp:]); l > 3 {
					continue
				}
				if n, _ := strconv.ParseInt(s[kp:], 10, 64); n > 255 {
					continue
				}
				if len(s[kp:]) > 1 && s[kp] == '0' {
					continue
				}
				result = append(result, fmt.Sprintf("%s.%s.%s.%s", s[:i], s[i:jp], s[jp:kp], s[kp:]))
			}
		}
	}
	return result
}
