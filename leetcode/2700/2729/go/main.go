package main

var fascinating = [10]bool{false, true, true, true, true, true, true, true, true, true}

func isFascinating(n int) bool {
	var (
		digits [10]bool
		ok     bool
	)

	for i := 1; i <= 3; i++ {
		if digits, ok = split(i*n, digits); !ok {
			return false
		}
	}

	return digits == fascinating
}

func split(n int, digits [10]bool) ([10]bool, bool) {
	for ; n > 0; n /= 10 {
		if d := n % 10; digits[d] || d == 0 {
			return digits, false
		} else {
			digits[d] = true
		}
	}

	return digits, true
}
