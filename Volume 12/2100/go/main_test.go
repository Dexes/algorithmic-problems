package main

import "testing"

type TestCase struct {
	Guests   []string
	Expected int
}

var TestCases = []TestCase{
	{[]string{"Ted+one", "Robin", "Barney"}, 600},
	{[]string{"Alice+one", "Bob+one", "ca", "Debora+one", "Enrique+one", "FRED+one"}, 1400},
}

func ReformatGuests(guests []string) [][]byte {
	result := make([][]byte, len(guests))
	for index, guest := range guests {
		result[index] = []byte(guest)
	}

	return result
}

func Test(t *testing.T) {
	for _, testcase := range TestCases {
		actual := GetResult(ReformatGuests(testcase.Guests))
		if actual != testcase.Expected {
			t.Error(
				"For", testcase.Guests,
				"expected", testcase.Expected,
				"got", actual,
			)
		}
	}
}
