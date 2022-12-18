package main

const typesCount = 5

var denominations = [typesCount]int{20, 50, 100, 200, 500}
var emptyResult = []int{-1}

type ATM struct {
	banknotes [typesCount]int
}

func Constructor() ATM {
	return ATM{banknotes: [typesCount]int{}}
}

func (x *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < typesCount; i++ {
		x.banknotes[i] += banknotesCount[i]
	}
}

func (x *ATM) Withdraw(amount int) []int {
	banknotes, count := x.banknotes, 0
	result := make([]int, len(denominations))
	for i := len(denominations) - 1; i >= 0 && amount > 0; i-- {
		if banknotes[i] == 0 || denominations[i] > amount {
			continue
		}

		if count = amount / denominations[i]; count > banknotes[i] {
			count = banknotes[i]
		}

		amount -= count * denominations[i]
		banknotes[i] -= count
		result[i] = count
	}

	if amount > 0 {
		return emptyResult
	}

	x.banknotes = banknotes
	return result
}
