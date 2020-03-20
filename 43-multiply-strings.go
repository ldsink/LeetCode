package main

func multiply(num1 string, num2 string) string {
	var product [300]int
	rs1 := []rune(num1)
	rs2 := []rune(num2)
	for i, j := 0, len(rs1); i < j; i++ {
		for k, l := 0, len(rs2); k < l; k++ {
			m := int(rs1[j-i-1] - '0')
			n := int(rs2[l-k-1] - '0')
			product[i+k] += m * n
		}
	}
	for i, j := 0, len(num1)+len(num2); i < j; i++ {
		product[i+1] += product[i] / 10
		product[i] = product[i] % 10
	}
	var result []rune
	for i := len(num1) + len(num2); i >= 0; i-- {
		if product[i] == 0 && i != 0 {
			continue
		}
		for j := i; j >= 0; j-- {
			result = append(result, rune('0'+product[j]))
		}
		break
	}
	return string(result)
}
