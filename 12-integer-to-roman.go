package main

type romanSymbol struct {
	Symbol string
	Value  int
}

func intToRoman(num int) string {
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
	var result string
	for _, symbol := range symbols {
		for ; num >= symbol.Value; num -= symbol.Value {
			result += symbol.Symbol
		}
	}
	return result
}
