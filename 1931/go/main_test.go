package main

import "testing"

type TestCase struct {
	Team     []int
	Expected int
}

var TestCases = []TestCase{
	{[]int{2, 5, 3, 4, 1, 9}, 1},
	{[]int{2, 3, 1, 4, 5}, 3},
	{[]int{5, 4, 3, 2, 1}, 2},
	{[]int{1}, 1},
	{[]int{9, 10, 6, 7, 8, 3, 4, 5, 2, 4, 11, 12}, 3},
}

func Test(t *testing.T) {
	for _, testcase := range TestCases {
		actual := GetResult(len(testcase.Team), testcase.Team)
		if actual != testcase.Expected {
			t.Error(
				"For", testcase.Team,
				"expected", testcase.Expected,
				"got", actual,
			)
		}
	}
}
