package main

const intA = int(rune('A'))
const intZ = int(rune('Z'))

func allLowerCase(rs []rune) bool {
	for i := 0; i < len(rs); i++ {
		if intA <= int(rs[i]) && int(rs[i]) <= intZ {
			return false
		}
	}
	return true
}

func allUpperCase(rs []rune) bool {
	for i := 0; i < len(rs); i++ {
		if !(intA <= int(rs[i]) && int(rs[i]) <= intZ) {
			return false
		}
	}
	return true
}

func detectCapitalUse(word string) bool {
	rs := []rune(word)
	if intA <= int(rs[0]) && int(rs[0]) <= intZ {
		return allLowerCase(rs[1:]) || allUpperCase(rs[1:])
	}
	return allLowerCase(rs[1:])
}
