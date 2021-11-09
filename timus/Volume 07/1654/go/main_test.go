package main

import "testing"

type TestCase struct {
	Data     string
	Expected string
}

var TestCases = []TestCase{
	{"wwstdaadierfflitzzz", "stierlitz"},
	{"", ""},
	{"a", "a"},
	{"ab", "ab"},
	{"dddd", ""},
	{"ddddd", "d"},
	{"abba", ""},
	{"abbbbac", "c"},
	{"wliisddsmmleeddw", ""},
	{"avcbbffcv", "a"},
	{"aaaxx", "a"},
}

func Test(t *testing.T) {
	for _, testcase := range TestCases {
		actual := GetResult([]byte(testcase.Data))
		if actual != testcase.Expected {
			t.Error(
				"For", testcase.Data,
				"expected", testcase.Expected,
				"got", actual,
			)
		}
	}
}
