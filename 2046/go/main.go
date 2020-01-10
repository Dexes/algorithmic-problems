package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetWord(subject string) (string, int) {
	length := len(subject)
	if length <= 10 {
		return subject, length
	}

	prevSpace := 0
	for i := 0; i <= 10; i++ {
		if subject[i] == 32 {
			prevSpace = i
		}
	}

	return subject[:prevSpace], prevSpace
}

func GetResult(data [3][4]string) string {
	border := "+----------+----------+----------+\n"
	result := border

	row := 0
	for row < 4 {
		incrementRow := true

		for column := 0; column < 3; column++ {
			word, length := GetWord(data[column][row])
			result += "|" + fmt.Sprintf("%-10s", word)
			if length == len(data[column][row]) {
				data[column][row] = ""
			} else {
				incrementRow = false
				data[column][row] = data[column][row][length+1:]
			}
		}

		result += "|\n"

		if incrementRow {
			row++
			result += border
		}
	}

	return result
}

func GetFirstRowSize(input []byte) int {
	length := len(input)
	result := 0

	for ; result < length; result++ {
		if input[result] == '\n' {
			break
		}
	}

	return result + 1
}

func GetPositionInTimetable(time []byte) (int, int) {
	length := len(time)
	number := int(time[length-1] - '1')

	if time[1] == 'u' {
		return 0, number
	}

	if time[1] == 'h' {
		return 1, number
	}

	return 2, number
}

func ParseInput(input []byte) [3][4]string {
	var size int

	result := [3][4]string{}
	length := len(input)
	pointer := GetFirstRowSize(input)

	for pointer < length {
		size = GetFirstRowSize(input[pointer:])
		subject := string(input[pointer : pointer+size-2])
		pointer += size

		size = GetFirstRowSize(input[pointer:])
		time := input[pointer : pointer+size-2]
		pointer += size

		x, y := GetPositionInTimetable(time)
		result[x][y] = subject
	}

	return result
}

func GetAllInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	println(GetResult(ParseInput(GetAllInput())))
}
