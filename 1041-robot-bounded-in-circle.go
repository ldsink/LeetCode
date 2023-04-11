package main

func isRobotBounded(instructions string) bool {
	var x, y, t int
	visited := make(map[int]bool)
	for i, j := 0, len(instructions)*len(instructions)+100; i <= j; i++ {
		k := i % len(instructions)
		p := (t << 60) | (x << 40) | (y << 20) | k
		if visited[p] {
			return true
		}
		visited[p] = true
		switch instructions[k] {
		case 'G':
			switch t {
			case 0:
				y++
			case 1:
				x++
			case 2:
				y--
			case 3:
				x--
			}
		case 'L':
			t = (t + 1) % 4
		case 'R':
			t = (t + 3) % 4
		}
	}
	return false
}
