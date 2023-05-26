package main

import (
	"bufio"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	if x := readInt(in); x > 2 && x&1 == 0 {
		_, _ = out.WriteString("YES")
	} else {
		_, _ = out.WriteString("NO")
	}

	_ = out.Flush()
}

func readInt(in *bufio.Reader) (result int) {
	for {
		if b, _ := in.ReadByte(); '0' <= b && b <= '9' {
			result = result*10 + int(b-'0')
			continue
		}

		return result
	}
}
