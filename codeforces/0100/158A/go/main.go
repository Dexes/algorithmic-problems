package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	n, k, min := readInt(in), readInt(in), 0
	for i := 0; i < k; i++ {
		if min = readInt(in); min == 0 {
			_, _ = out.WriteString(strconv.Itoa(i))
			_ = out.Flush()
			return
		}
	}

	for i := n - k; i > 0; i-- {
		if x := readInt(in); x == min {
			continue
		}

		_, _ = out.WriteString(strconv.Itoa(n - i))
		_ = out.Flush()
		return
	}

	_, _ = out.WriteString(strconv.Itoa(n))
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
