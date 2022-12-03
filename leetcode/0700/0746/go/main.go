package main

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i]+cost[i-2], cost[i]+cost[i-1])
	}

	return min(cost[len(cost)-2], cost[len(cost)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
