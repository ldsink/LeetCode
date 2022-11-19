package main

func largestAltitude(gain []int) int {
	var altitude, result int
	for i := 0; i < len(gain); i++ {
		altitude += gain[i]
		if result < altitude {
			result = altitude
		}
	}
	return result
}
