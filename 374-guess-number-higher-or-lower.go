package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	start := 1
	for start < n {
		middle := (start + n) >> 1
		if r := guess(middle); r == 0 {
			return middle
		} else if r == -1 {
			n = middle - 1
		} else {
			start = middle + 1
		}
	}
	return n
}
