package main

import (
	"bufio"
	"os"
)

const SymbolTypeBigLetter = 1
const SymbolTypeSmallLetter = 2
const SymbolTypeSentenceEnd = 3
const SymbolTypeOther = 4

func GetSymbolType(symbol byte) int {
	if 'a' <= symbol && symbol <= 'z' {
		return SymbolTypeSmallLetter
	}

	if 'A' <= symbol && symbol <= 'Z' {
		return SymbolTypeBigLetter
	}

	if symbol == '.' || symbol == '?' || symbol == '!' {
		return SymbolTypeSentenceEnd
	}

	return SymbolTypeOther
}

func ProcessSentence(text []byte) (int, int) {
	length := len(text)
	firstLetterPassed := false
	previousIsLetter := false
	errorsCount := 0
	size := 0

	for size = 0; size < length; size++ {
		symbolType := GetSymbolType(text[size])
		if symbolType == SymbolTypeSentenceEnd {
			break
		}

		if symbolType == SymbolTypeBigLetter && previousIsLetter {
			errorsCount++
		} else if symbolType == SymbolTypeSmallLetter && !firstLetterPassed {
			errorsCount++
		}

		previousIsLetter = symbolType == SymbolTypeBigLetter || symbolType == SymbolTypeSmallLetter
		firstLetterPassed = firstLetterPassed || previousIsLetter
	}

	return errorsCount, size
}

func GetResult(text []byte) (count int) {
	pointer := 0
	length := len(text)
	for pointer < length {
		errorsCount, size := ProcessSentence(text[pointer:])
		count += errorsCount
		pointer += size + 1
	}

	return
}

func GetAllInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	println(GetResult(GetAllInput()))
}
