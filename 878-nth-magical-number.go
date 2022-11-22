package main

import (
	"math/big"
)

func nthMagicalNumber(n int, a int, b int) int {
	// 最大公约数 gcd 和 最小公倍数 lcm
	var gcd, lcm func(int, int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	lcm = func(a, b int) int {
		return a * b / gcd(a, b)
	}
	// 二分查找
	bigA, bigB, bigP, bigN := big.NewInt(int64(a)), big.NewInt(int64(b)), big.NewInt(int64(lcm(a, b))), big.NewInt(int64(n))
	one := big.NewInt(1)
	start := big.NewInt(0)
	end, _ := new(big.Int).SetString("40000000000000", 10)
	var middle, aN, bN, pN, idx big.Int
	for start.Cmp(end) != 0 {
		middle.Add(start, end)
		middle.Rsh(&middle, 1)
		aN.Div(&middle, bigA)
		bN.Div(&middle, bigB)
		pN.Div(&middle, bigP)
		idx.Add(&aN, &bN).Sub(&idx, &pN)
		if idx.Cmp(bigN) < 0 {
			start.Add(&middle, one)
		} else {
			end.Set(&middle)
		}
	}
	start.Mod(start, big.NewInt(1e9+7))
	return int(start.Int64())
}
