package main

func accountBalanceAfterPurchase(purchaseAmount int) int {
	mod := purchaseAmount % 10
	if mod > 4 {
		purchaseAmount += 10
	}

	return 100 - purchaseAmount + mod
}
