package main

func gcdOfStrings(str1 string, str2 string) string {
	for cpl := commonPrefixLength(str1, str2); cpl > 0; cpl-- {
		if len(str1)%cpl > 0 || len(str2)%cpl > 0 {
			continue
		}

		if isDivisor(str1, cpl) && isDivisor(str2, cpl) {
			return str1[:cpl]
		}
	}

	return ""
}

func isDivisor(s string, length int) bool {
	for i := 0; i < length; i++ {
		for j := i + length; j < len(s); j += length {
			if s[i] != s[j] {
				return false
			}
		}
	}

	return true
}

func commonPrefixLength(s1, s2 string) int {
	i := 0
	for ; i < len(s1) && i < len(s2) && s1[i] == s2[i]; i++ {
	}

	return i
}
