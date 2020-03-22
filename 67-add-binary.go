package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ar := []rune(a)
	br := []rune(b)
	for i, j := 0, len(ar)-1; i <= j; i, j = i+1, j-1 {
		ar[i], ar[j] = ar[j], ar[i]
	}
	for i, j := 0, len(br)-1; i <= j; i, j = i+1, j-1 {
		br[i], br[j] = br[j], br[i]
	}

	var result []rune
	var bit int
	var up bool
	for ; bit < len(br); bit++ {
		if ar[bit] == '0' && br[bit] == '0' {
			if up {
				result = append(result, '1')
				up = false
			} else {
				result = append(result, '0')
			}
		} else if ar[bit] == '1' && br[bit] == '1' {
			if up {
				result = append(result, '1')
			} else {
				result = append(result, '0')
				up = true
			}
		} else {
			if up {
				result = append(result, '0')
			} else {
				result = append(result, '1')
			}
		}
	}
	for ; bit < len(ar); bit++ {
		if !up {
			result = append(result, ar[bit])
		} else if ar[bit] == '1' {
			result = append(result, '0')
		} else {
			result = append(result, '1')
			up = false
		}
	}
	if up {
		result = append(result, '1')
		up = false
	}
	for i, j := 0, len(result)-1; i <= j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
