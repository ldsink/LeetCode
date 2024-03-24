package main

func distinctIntegers(n int) int {
	set := make(map[int]bool)
	set[n] = true
	for i := 1; i <= n; i++ {
		newNums := make(map[int]bool)
		for x := range set {
			for i := 1; i <= n; i++ {
				if set[i] || newNums[i] {
					continue
				}
				if x%i == 1 {
					newNums[i] = true
				}
			}
		}
		for i := range newNums {
			set[i] = true
		}
		if len(set) == n {
			break
		}
	}
	return len(set)
}
