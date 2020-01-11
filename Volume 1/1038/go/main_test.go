package main

import "testing"

type TestCase struct {
	Sentence string
	Expected int
}

var TestCases = []TestCase{
	{"This sentence iz correkt! -It Has,No mista;.Kes et oll.\nBut there are two BIG mistakes in this one!\nand here is one more.", 3},
	{"But there are two BIG mistakes in this one!\nand here is one more.", 3},
	{"and here is one more.", 1},
	{" and here is one more.", 1},
	{"But there are two BIG mistakes in this one!", 2},
	{"but there are two BIG mistakes in this one!", 3},
	{"This sentence iz correkt! -It Has,No mista;.Kes et oll.", 0},
	{"This sentence iz correkt", 0},
}

func Test(t *testing.T) {
	for _, testcase := range TestCases {
		actual := GetResult([]byte(testcase.Sentence))
		if actual != testcase.Expected {
			t.Error(
				"For", testcase.Sentence,
				"expected", testcase.Expected,
				"got", actual,
			)
		}
	}
}
