package main

func myAtoi(str string) int {
	const plusSign rune = '+'
	const minusSign rune = '-'
	const whitespaceSign rune = ' '
	const IntMax int64 = 1<<31 - 1
	const IntMin int64 = -(1 << 31)
	numSet := make(map[rune]int)
	numRS := []rune("0123456789")
	for n, r := range numRS {
		numSet[r] = n
	}

	var signFlag, numberFlag int
	var numbers []int
	rs := []rune(str)
	for _, r := range rs {
		if r == whitespaceSign {
			if signFlag != 0 || numberFlag != 0 {
				break
			}
			continue
		} else if r == plusSign || r == minusSign {
			if signFlag != 0 || numberFlag != 0 {
				break
			}
			if r == plusSign {
				signFlag = 1
			} else {
				signFlag = -1
			}
			continue
		} else if number, ok := numSet[r]; ok {
			numberFlag = 1
			numbers = append(numbers, number)
			continue
		} else {
			break
		}
	}
	if len(numbers) == 0 {
		return 0
	}

	var zeroOffset int
	for _, r := range numbers {
		if r == 0 {
			zeroOffset++
		} else {
			break
		}
	}
	numbers = numbers[0+zeroOffset : len(numbers)]

	if signFlag == 0 {
		signFlag = 1
	}

	if len(numbers) > 13 {
		if signFlag == 1 {
			return int(IntMax)
		} else {
			return int(IntMin)
		}
	}

	var result int64
	for i, j := 0, int64(1); i < len(numbers); i, j = i+1, j*10 {
		result += j * int64(numbers[len(numbers)-i-1])
	}
	result *= int64(signFlag)
	if result > IntMax {
		return int(IntMax)
	} else if result < IntMin {
		return int(IntMin)
	}
	return int(result)
}
