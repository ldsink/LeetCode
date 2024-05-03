package main

func average(salary []int) float64 {
	var sum, max, min int
	max, min = salary[0], salary[0]
	for _, s := range salary {
		sum += s
		if min > s {
			min = s
		}
		if max < s {
			max = s
		}
	}
	return float64(sum-max-min) / float64(len(salary)-2)
}
