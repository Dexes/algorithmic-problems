package main

func countSeniors(details []string) (result int) {
	for _, d := range details {
		if (d[11]-'0')*10+(d[12]-'0') > 60 {
			result++
		}
	}

	return result
}
