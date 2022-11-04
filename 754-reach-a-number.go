package main

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	numMoves := 1
	for ; target > 0; numMoves++ {
		target -= numMoves
	}
	numMoves--
	if target%2 == 0 {
		return numMoves
	}
	return numMoves + 1 + numMoves%2
}
