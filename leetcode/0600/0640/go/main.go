package main

import "strconv"

func main() {
	println(solveEquation("-x=-1"))
}

func solveEquation(equation string) string {
	x, num, invert := 0, 0, false
	sign, value, isX := 0, 0, false

	for i := 0; i < len(equation); {
		sign, value, isX, i = getNum(equation, i)

		if isX {
			if invert {
				sign = -sign
			}

			x += sign * value
		} else {
			if !invert {
				sign = -sign
			}

			num += sign * value
		}

		if i < len(equation) && equation[i] == '=' {
			invert, i = true, i+1
		}
	}

	if x == 0 {
		if num == 0 {
			return "Infinite solutions"
		}

		return "No solution"
	}

	return "x=" + strconv.Itoa(num/x)
}

func getNum(s string, i int) (int, int, bool, int) {
	sign, num := 1, 0

	if s[i] == '-' {
		sign, i = -1, i+1
	} else if s[i] == '+' {
		i++
	}

	if s[i] == 'x' {
		return sign, 1, true, i + 1
	}

	for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		num = num*10 + int(s[i]-'0')
	}

	if i < len(s) && s[i] == 'x' {
		return sign, num, true, i + 1
	}

	return sign, num, false, i
}
