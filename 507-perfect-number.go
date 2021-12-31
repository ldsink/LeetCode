package main

func checkPerfectNumber(num int) bool {
	if num <= 3 {
		return false
	}
	sum := 1
	i := 2
	for i = 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			sum += num / i
			if sum > num {
				return false
			}
		}
	}
	if (i-1)*(i-1) == num {
		sum -= i - 1
	}
	return sum == num
}
