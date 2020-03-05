package main

var NumRunes = map[rune][]rune{
	'2': []rune("abc"),
	'3': []rune("def"),
	'4': []rune("ghi"),
	'5': []rune("jkl"),
	'6': []rune("mno"),
	'7': []rune("pqrs"),
	'8': []rune("tuv"),
	'9': []rune("wxyz"),
}

func _processCombine(rs []rune, prefix []rune) []string {
	if len(rs) == 0 {
		return []string{string(prefix)}
	}
	var result []string
	for _, r := range NumRunes[rs[0]] {
		current := _processCombine(rs[1:], append(prefix, r))
		result = append(result, current...)
	}
	return result
}

func letterCombinations(digits string) []string {
	var result []string
	rs := []rune(digits)
	if len(rs) == 0 {
		return result
	}
	var prefix []rune
	current := _processCombine(rs, prefix)
	result = append(result, current...)
	return result
}
