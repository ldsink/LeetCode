package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	press, result := releaseTimes[0], keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		j := releaseTimes[i] - releaseTimes[i-1]
		if press < j || (press == j && result < keysPressed[i]) {
			press, result = j, keysPressed[i]
		}
	}
	return result
}
