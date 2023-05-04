package main

const (
	ResultRadiant = "Radiant"
	ResultDire    = "Dire"
)

func predictPartyVictory(senate string) string {
	bytes, balance, wIndex := []byte(senate), 0, 0
	for abs(balance) < len(bytes) {
		for i := 0; i < len(bytes); i++ {
			x := toInt(bytes[i])
			if balance == 0 || (balance < 0) == (x < 0) {
				bytes[wIndex] = bytes[i]
				wIndex++
			}

			balance += x
		}

		bytes, wIndex = bytes[:wIndex], 0
	}

	if balance < 0 {
		return ResultRadiant
	}

	return ResultDire
}

func toInt(b byte) int {
	if b == 'R' {
		return -1
	}

	return 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
