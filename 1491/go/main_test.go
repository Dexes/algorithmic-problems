package main

import (
	"encoding/json"
	"testing"
)

type TestCase struct {
	JSONData     string
	JSONExpected string
}

var TestCases = []TestCase{
	{"[[2,3,2],[2,4,1],[3,4,1],[1,2,1],[1,1,1]]", "[2,4,4,2]"},
	{"[[1,7,1],[2,3,4],[3,5,3],[1,2,1],[5,7,4],[2,4,10],[3,4,2],[1,6,3]]", "[5,19,23,19,11,8,5]"},
}

func Test(t *testing.T) {
	for _, testcase := range TestCases {
		var data [][]int
		json.Unmarshal([]byte(testcase.JSONData), &data)

		actual, _ := json.Marshal(GetResult(len(data)-1, data))
		if string(actual) != testcase.JSONExpected {
			t.Error(
				"For", testcase.JSONData,
				"expected", testcase.JSONExpected,
				"got", string(actual),
			)
		}
	}
}
