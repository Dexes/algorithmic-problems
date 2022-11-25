package main

type Logger struct {
	messages map[string]int
}

func Constructor() Logger {
	return Logger{messages: make(map[string]int)}
}

func (x *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if prev, ok := x.messages[message]; ok && timestamp-prev < 10 {
		return false
	}

	x.messages[message] = timestamp
	return true
}
