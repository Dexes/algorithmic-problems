package main

import (
	"bufio"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	var row, col byte
	for {
		switch b, _ := in.ReadByte(); {
		case b == '1':
			_ = out.WriteByte(count(col) + count(row) + '0')
			_ = out.Flush()
			return
		case b == '0':
			_, _ = in.ReadByte()
			col++
		default:
			col, row = 0, row+1
		}
	}
}

func count(x byte) byte {
	if x > 2 {
		return x - 2
	}

	return 2 - x
}
