package main

func playGame(friends []int, k, index int) ([]int, int) {
	next := (index + k - 1) % len(friends)
	friends = append(friends[:next], friends[next+1:]...)
	if next == len(friends) {
		next = 0
	}
	return friends, next
}

func findTheWinner(n int, k int) int {
	if k == 1 || n == 1 {
		return n
	}
	friends := make([]int, n)
	for i := 0; i < n; i++ {
		friends[i] = i + 1
	}
	for i := 0; len(friends) > 1; {
		friends, i = playGame(friends, k, i)
	}
	return friends[0]
}
