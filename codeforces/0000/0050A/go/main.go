package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	x, y := readInt(in), readInt(in)

	_, _ = out.WriteString(strconv.Itoa(x * y / 2))
	_ = out.Flush()
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
