package main

import "testing"

type TestCase struct {
	Heap     []int
	Expected int
}

var TestCases = []TestCase{
	{[]int{5, 8, 13, 27, 14}, 3},
	{[]int{5, 5, 4, 3, 3}, 0},
	{[]int{6, 7, 9, 13, 18, 24, 31, 50}, 0},
	{[]int{1, 2, 3, 4, 100, 100}, 0},
	{[]int{15, 16, 30, 2}, 1},
	{[]int{5, 2, 3, 2, 3, 2}, 1},
	{[]int{1}, 1},
}

func Test(t *testing.T) {
	for _, testcase := range TestCases {
		actual := GetResult(len(testcase.Heap), testcase.Heap)
		if actual != testcase.Expected {
			t.Error(
				"For", testcase.Heap,
				"expected", testcase.Expected,
				"got", actual,
			)
		}
	}
}
