package main

func calculate(s string) int {
	operations := readOperations(s)
	return operations.Calculate()
}

type Operator byte

func (operator Operator) GetPriority() byte {
	if operator == '+' || operator == '-' {
		return 2
	}

	return 1
}

func (operator Operator) Apply(a, b int) int {
	switch operator {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		return a / b
	}
}

type OperationNode struct {
	Number   int
	Operator Operator
	Next     *OperationNode
}

func (first *OperationNode) Calculate() int {
	for priority := byte(1); priority <= 2; priority++ {
		node := first
		for node.Next != nil {
			if node.Operator.GetPriority() != priority {
				node = node.Next
				continue
			}

			node.Number = node.Operator.Apply(node.Number, node.Next.Number)
			node.Operator = node.Next.Operator
			node.Next = node.Next.Next
		}
	}

	return first.Number
}

func readOperations(s string) *OperationNode {
	var (
		first, current *OperationNode
		number, pos    int
		operator       Operator
	)

	for {
		number, pos = readInt(s, pos)
		operator, pos = readOperator(s, pos)
		node := &OperationNode{Number: number, Operator: operator}

		if first == nil {
			first, current = node, node
		} else {
			current.Next, current = node, node
		}

		if pos == len(s) {
			return first
		}
	}
}

func readOperator(s string, pos int) (Operator, int) {
	pos = skipSpaces(s, pos)
	if pos == len(s) {
		return 0, pos
	}

	return Operator(s[pos]), pos + 1
}

func readInt(s string, pos int) (int, int) {
	result := 0
	for pos = skipSpaces(s, pos); pos < len(s) && isDigit(s[pos]); pos++ {
		result = result*10 + int(s[pos]-'0')
	}

	return result, pos
}

func skipSpaces(s string, pos int) int {
	for ; pos < len(s) && s[pos] == ' '; pos++ {
	}

	return pos
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
