package main

type romanSymbol struct {
	Symbol string
	Value  int
}

func romanToInt(s string) int {
	symbols := []romanSymbol{
		{Symbol: "M", Value: 1000},
		{Symbol: "CM", Value: 900},
		{Symbol: "D", Value: 500},
		{Symbol: "CD", Value: 400},
		{Symbol: "C", Value: 100},
		{Symbol: "XC", Value: 90},
		{Symbol: "L", Value: 50},
		{Symbol: "XL", Value: 40},
		{Symbol: "X", Value: 10},
		{Symbol: "X", Value: 10},
		{Symbol: "IX", Value: 9},
		{Symbol: "V", Value: 5},
		{Symbol: "IV", Value: 4},
		{Symbol: "I", Value: 1},
	}
	var result int
	for _, symbol := range symbols {
		for ; len(s) >= len(symbol.Symbol) && s[:len(symbol.Symbol)] == symbol.Symbol; s = s[len(symbol.Symbol):] {
			result += symbol.Value
		}
	}
	return result
}
