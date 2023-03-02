package main

import "strings"

type Codec struct {
}

func (codec Codec) Encode(s []string) string {
	sb := strings.Builder{}
	sb.WriteByte(byte(len(s)))

	for i := 0; i < len(s); i++ {
		sb.WriteByte(byte(len(s[i])))
		sb.WriteString(s[i])
	}

	return sb.String()
}

func (codec Codec) Decode(s string) []string {
	result := make([]string, 0, s[0])
	for index, next := 1, 0; index < len(s); index = next {
		next = index + int(s[index]) + 1
		result = append(result, s[index+1:next])
	}

	return result
}
