package main

import (
	"bufio"
	"os"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

func StackToString(stack *Stack) string {
	length := stack.Len()
	result := make([]byte, length)

	for i := 1; i <= length; i++ {
		result[length-i] = stack.Pop().(byte)
	}

	return string(result)
}

func GetResult(word []byte) string {
	stack := NewStack()
	length := len(word)

	for i := 0; i < length; i++ {
		if stack.Peek() == word[i] {
			stack.Pop()
		} else {
			stack.Push(word[i])
		}
	}

	return StackToString(stack)
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	bytes, _ := scanner.ReadBytes(0)

	println(GetResult(bytes))
}
