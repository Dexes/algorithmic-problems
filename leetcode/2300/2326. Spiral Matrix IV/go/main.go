package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(rows, cols int, head *ListNode) [][]int {
	result, emptyCells := makeMatrix(rows, cols), rows*cols
	for x := 0; ; x++ {
		firstRow, secondRow, firstCol, secondCol := x, rows-x-1, x, cols-x-1

		for col := firstCol; col <= secondCol; col++ {
			result[firstRow][col], head = value(head)
		}

		if emptyCells -= secondCol - firstCol + 1; emptyCells == 0 {
			return result
		}

		for row := firstRow + 1; row <= secondRow; row++ {
			result[row][secondCol], head = value(head)
		}

		if emptyCells -= secondRow - firstRow; emptyCells == 0 {
			return result
		}

		for col := secondCol - 1; col >= firstCol; col-- {
			result[secondRow][col], head = value(head)
		}

		if emptyCells -= secondCol - firstCol; emptyCells == 0 {
			return result
		}

		for row := secondRow - 1; row > firstRow; row-- {
			result[row][firstCol], head = value(head)
		}

		if emptyCells -= secondRow - firstRow - 1; emptyCells == 0 {
			return result
		}
	}
}

func value(node *ListNode) (int, *ListNode) {
	if node == nil {
		return -1, nil
	}

	return node.Val, node.Next
}

func makeMatrix(rows, cols int) [][]int {
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}

	return result
}
