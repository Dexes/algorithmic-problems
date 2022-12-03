package main

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	bytes := make([]byte, 4)

	return func(buf []byte, n int) int {
		for buf = buf[:0]; ; {
			switch x := read4(bytes); {
			case len(buf)+x >= n:
				buf = append(buf, bytes[:n-len(buf)]...)
				return len(buf)

			case x < 4:
				buf = append(buf, bytes[:x]...)
				return len(buf)

			default:
				buf = append(buf, bytes...)
			}
		}
	}
}
