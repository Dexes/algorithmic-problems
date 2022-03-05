package main

func totalMoney(n int) int {
	weeks, days := n/7, n%7
	return (7*weeks*(weeks+7) + days*(weeks*2+days+1)) / 2
}
