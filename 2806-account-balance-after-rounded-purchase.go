package main

func accountBalanceAfterPurchase(purchaseAmount int) int {
	var cost int
	delta := purchaseAmount % 10
	if delta < 5 {
		cost = purchaseAmount - delta
	} else {
		cost = purchaseAmount + 10 - delta
	}
	return 100 - cost
}
