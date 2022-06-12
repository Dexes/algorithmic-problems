package main

const (
	lowercaseCheck   = 1 << iota
	uppercaseCheck   = 1 << iota
	digitCheck       = 1 << iota
	specialCharCheck = 1 << iota
	validMask        = 1<<iota - 1
)

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	mask, isValid := 0, false
	for i := 0; i < len(password); i++ {
		if i > 0 && password[i-1] == password[i] {
			return false
		}

		switch x := password[i]; {
		case isValid:
			continue
		case 'a' <= x && x <= 'z':
			mask |= lowercaseCheck
		case 'A' <= x && x <= 'Z':
			mask |= uppercaseCheck
		case '0' <= x && x <= '9':
			mask |= digitCheck
		default:
			mask |= specialCharCheck
		}

		isValid = mask == validMask
	}

	return isValid
}
