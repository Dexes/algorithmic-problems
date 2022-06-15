package main

import (
	"sort"
)

type transaction struct {
	info    string
	amount  int
	time    int
	city    string
	invalid bool
}

func invalidTransactions(transactions []string) []string {
	wIndex := 0

	for _, uTransactions := range parseTransactions(transactions) {
		markInvalid(uTransactions)

		for i := 0; i < len(uTransactions); i++ {
			if uTransactions[i].invalid {
				transactions[wIndex] = uTransactions[i].info
				wIndex++
			}
		}
	}

	return transactions[:wIndex]
}

func markInvalid(transactions []transaction) {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].time < transactions[j].time
	})

	for i := 0; i < len(transactions); i++ {
		if transactions[i].amount > 1000 {
			transactions[i].invalid = true
		}

		time := transactions[i].time + 60
		for j := i + 1; j < len(transactions) && transactions[j].time <= time; j++ {
			if transactions[j].city != transactions[i].city {
				transactions[i].invalid, transactions[j].invalid = true, true
			}
		}
	}
}

func parseTransactions(transactions []string) map[string][]transaction {
	var time, amount, index int

	result := make(map[string][]transaction)
	for _, t := range transactions {
		name := t[:search(t, 0, ',')]
		time, index = toInt(t, len(name)+1)
		amount, index = toInt(t, index+1)
		city := t[index+1:]

		if result[name] == nil {
			result[name] = make([]transaction, 0, 10)
		}

		result[name] = append(result[name], transaction{info: t, time: time, amount: amount, city: city})
	}

	return result
}

func search(s string, index int, b byte) int {
	for ; index < len(s); index++ {
		if s[index] == b {
			return index
		}
	}

	return -1
}

func toInt(s string, index int) (int, int) {
	result := 0
	for ; '0' <= s[index] && s[index] <= '9'; index++ {
		result = result*10 + int(s[index]-'0')
	}

	return result, index
}
