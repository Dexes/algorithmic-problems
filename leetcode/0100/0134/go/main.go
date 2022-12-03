package main

func canCompleteCircuit(gas []int, cost []int) int {
	result, sum, tank := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		sum, tank = sum+diff, tank+diff
		if tank < 0 {
			result = i + 1
			tank = 0
		}
	}

	if sum < 0 {
		return -1
	}

	return result
}
