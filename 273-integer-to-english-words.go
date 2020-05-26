package main

import (
	"strconv"
	"strings"
)

var unitsMap, tensMap map[rune]string
var tenMap map[string]string

func numStrToWords(s string) string {
	var result string
	if len(s) > 9 {
		result = numStrToWords(s[:len(s)-9])
		if result != "" {
			result += " Billion "
		}
		result += numStrToWords(s[len(s)-9:])
		return strings.TrimSpace(result)
	} else if len(s) > 6 {
		result = numStrToWords(s[:len(s)-6])
		if result != "" {
			result += " Million "
		}
		result += numStrToWords(s[len(s)-6:])
		return strings.TrimSpace(result)
	} else if len(s) > 3 {
		result = numStrToWords(s[:len(s)-3])
		if result != "" {
			result += " Thousand "
		}
		result += numStrToWords(s[len(s)-3:])
		return strings.TrimSpace(result)
	} else if len(s) == 3 {
		result := numStrToWords(s[1:])
		if s[0] != '0' {
			result = unitsMap[rune(s[0])] + " Hundred " + result
		}
		return strings.TrimSpace(result)
	} else if len(s) == 2 {
		if s[0] == '0' {
			return numStrToWords(s[1:])
		} else if s[0] == '1' {
			return tenMap[s]
		}
		result := tensMap[rune(s[0])] + " " + numStrToWords(s[1:])
		return strings.TrimSpace(result)
	}
	if r, ok := unitsMap[rune(s[0])]; ok {
		return r
	}
	return ""
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	unitsMap = map[rune]string{
		'1': "One",
		'2': "Two",
		'3': "Three",
		'4': "Four",
		'5': "Five",
		'6': "Six",
		'7': "Seven",
		'8': "Eight",
		'9': "Nine",
	}
	tensMap = map[rune]string{
		'2': "Twenty",
		'3': "Thirty",
		'4': "Forty",
		'5': "Fifty",
		'6': "Sixty",
		'7': "Seventy",
		'8': "Eighty",
		'9': "Ninety",
	}
	tenMap = map[string]string{
		"10": "Ten",
		"11": "Eleven",
		"12": "Twelve",
		"13": "Thirteen",
		"14": "Fourteen",
		"15": "Fifteen",
		"16": "Sixteen",
		"17": "Seventeen",
		"18": "Eighteen",
		"19": "Nineteen",
	}
	s := strconv.Itoa(num)
	return numStrToWords(s)
}
