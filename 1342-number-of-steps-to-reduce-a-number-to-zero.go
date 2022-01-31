package main

func numberOfSteps(num int) int {
	var steps int
	for ; num > 0; steps++ {
		if num&1 == 1 {
			num ^= 1
		} else {
			num >>= 1
		}
	}
	return steps
}
