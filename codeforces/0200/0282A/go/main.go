package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	result := 0
	for i := readInt(in); i > 0; i-- {
		if line := readLine(in); line[1] == '-' {
			result--
		} else {
			result++
		}
	}

	_, _ = out.WriteString(strconv.Itoa(result))
	_ = out.Flush()
}

var line = make([]byte, 4)

func readLine(in *bufio.Reader) []byte {
	for i := 0; ; i++ {
		if line[i], _ = in.ReadByte(); line[i] == '\r' {
			_, _ = in.ReadByte()
			return line[:i]
		}
	}
}

func readInt(in *bufio.Reader) (result int) {
	for {
		switch b, _ := in.ReadByte(); {
		case '0' <= b && b <= '9':
			result = result*10 + int(b-'0')
		case b == '\r':
			_, _ = in.ReadByte()
			return result
		default:
			return result
		}
	}
}
