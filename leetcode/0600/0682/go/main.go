package main

import (
	"strconv"
)

func calPoints(ops []string) int {
	stack, stackIndex := make([]int, 1000), 0
	result := 0
	for i := 0; i < len(ops); i++ {
		switch ops[i][0] {
		case 'C':
			stackIndex--
			result -= stack[stackIndex]
		case 'D':
			stack[stackIndex] = stack[stackIndex-1] << 1
			result += stack[stackIndex]
			stackIndex++
		case '+':
			stack[stackIndex] = stack[stackIndex-1] + stack[stackIndex-2]
			result += stack[stackIndex]
			stackIndex++
		default:
			stack[stackIndex], _ = strconv.Atoi(ops[i])
			result += stack[stackIndex]
			stackIndex++
		}
	}

	return result
}
