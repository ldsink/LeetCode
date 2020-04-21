package main

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	minHP := [2][]int{make([]int, n), make([]int, n)}

	// initial last line
	if hp := dungeon[m-1][n-1]; hp <= 0 {
		minHP[(m-1)%2][n-1] = 1 - dungeon[m-1][n-1]
	} else {
		minHP[(m-1)%2][n-1] = 1
	}
	for i, j := n-2, m-1; i >= 0; i-- {
		k := j % 2
		if hp := dungeon[m-1][i]; hp <= 0 {
			minHP[k][i] = minHP[k][i+1] - hp
		} else if hp-minHP[k][i+1] >= 0 {
			minHP[k][i] = 1
		} else {
			minHP[k][i] = minHP[k][i+1] - hp
		}
	}

	// dynamic search
	for j := m - 2; j >= 0; j-- {
		k := j % 2
		if hp := dungeon[j][n-1]; hp <= 0 {
			minHP[k][n-1] = minHP[k^1][n-1] - dungeon[j][n-1]
		} else if hp-minHP[k^1][n-1] >= 0 {
			minHP[k][n-1] = 1
		} else {
			minHP[k][n-1] = minHP[k^1][n-1] - hp
		}
		for i := n - 2; i >= 0; i-- {
			hp := dungeon[j][i]
			var rightMin, downMin int

			// go right
			if hp <= 0 {
				rightMin = minHP[k][i+1] - hp
			} else if hp-minHP[k][i+1] >= 0 {
				rightMin = 1
			} else {
				rightMin = minHP[k][i+1] - hp
			}

			// go down
			if hp <= 0 {
				downMin = minHP[k^1][i] - hp
			} else if hp-minHP[k^1][i] >= 0 {
				downMin = 1
			} else {
				downMin = minHP[k^1][i] - hp
			}

			// get min
			if rightMin > downMin {
				minHP[k][i] = downMin
			} else {
				minHP[k][i] = rightMin
			}
		}
	}

	return minHP[0][0]
}
