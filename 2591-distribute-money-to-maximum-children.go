package main

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	money -= children
	if money > children*7 {
		return children - 1
	} else if money == children*7 {
		return children
	} else if money == children*7-4 {
		if children >= 2 {
			return children - 2
		}
		return -1
	}
	return money / 7
}
