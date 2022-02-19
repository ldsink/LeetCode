package main

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			if i+1 < len(bits) {
				bits[i+1] = 1
				i++
			}
		}
	}
	return bits[len(bits)-1] == 0
}
