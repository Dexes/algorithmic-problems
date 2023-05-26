package main

import (
	"bufio"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	for i := readInt(in); i > 0; i-- {
		line := readLine(in)
		letters := len(line) - 2

		switch {
		case letters < 9:
			_, _ = out.Write(line)
			_ = out.WriteByte('\n')
		case letters < 10:
			line[1] = byte(letters + '0')
			line[2] = line[len(line)-1]
			_, _ = out.Write(line[:3])
			_ = out.WriteByte('\n')
		case letters < 100:
			line[1] = byte(letters/10 + '0')
			line[2] = byte(letters%10 + '0')
			line[3] = line[len(line)-1]
			_, _ = out.Write(line[:4])
			_ = out.WriteByte('\n')
		default:
			line[1] = '1'
			line[2] = '0'
			line[3] = '0'
			line[4] = line[len(line)-1]
			_, _ = out.Write(line[:5])
			_ = out.WriteByte('\n')
		}
	}

	_ = out.Flush()
}

var line = make([]byte, 101)

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
