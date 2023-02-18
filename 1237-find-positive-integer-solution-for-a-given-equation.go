package main

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *              Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	var result [][]int
	var start, end int
	for x := 1; x <= 1000; x++ {
		if z < customFunction(x, 1) || customFunction(x, 1000) < z {
			continue
		}
		for start, end = 1, 1000; start < end; {
			middle := (start + end) >> 1
			if z <= customFunction(x, middle) {
				end = middle
			} else {
				start = middle + 1
			}
		}
		if customFunction(x, start) == z {
			result = append(result, []int{x, start})
		}
	}
	return result
}
