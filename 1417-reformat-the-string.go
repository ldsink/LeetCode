package main

func reformat(s string) string {
	var alphas, digits []rune
	for _, r := range []rune(s) {
		if '0' <= r && r <= '9' {
			digits = append(digits, r)
		} else {
			alphas = append(alphas, r)
		}
	}

	abs := func(n int) int {
		if n >= 0 {
			return n
		}
		return -n
	}
	if abs(len(alphas)-len(digits)) > 1 {
		return ""
	}

	var result []rune
	if len(alphas) < len(digits) {
		alphas, digits = digits, alphas
	}
	for i := 0; i < len(digits); i++ {
		result = append(result, alphas[i])
		result = append(result, digits[i])
	}
	if len(alphas) > len(digits) {
		result = append(result, alphas[len(alphas)-1])
	}
	return string(result)
}
