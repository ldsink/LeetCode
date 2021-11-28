package main

const Length = 26
const RuneA = 'a'

func isSame(a, b []int) bool {
	for i := 0; i < Length; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	var result []int
	if len(s) < len(p) {
		return result
	}
	// 统计 p 的字符情况
	dst := make([]int, Length)
	for _, r := range []rune(p) {
		dst[r-RuneA]++
	}
	// 遍历比较
	src := make([]int, Length)
	rs := []rune(s)
	for i := 0; i < len(p); i++ {
		src[rs[i]-RuneA]++
	}
	for i, j := 0, len(s)-len(p); i <= j; i++ {
		if isSame(src, dst) {
			result = append(result, i)
		}
		src[rs[i]-RuneA]--
		if i+len(p) < len(s) {
			src[rs[i+len(p)]-RuneA]++
		}
	}
	return result
}
