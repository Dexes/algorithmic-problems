package main

func numUniqueEmails(emails []string) int {
	keys := make(map[string]bool)
	for i := 0; i < len(emails); i++ {
		key := emailToKey(emails[i])
		keys[key] = true
	}

	return len(keys)
}

func emailToKey(email string) string {
	result := make([]byte, 0, len(email))
	i := 0

	for ; email[i] != '@' && email[i] != '+'; i++ {
		if email[i] == '.' {
			continue
		}

		result = append(result, email[i])
	}

	for ; email[i] != '@'; i++ {
	}

	for ; i < len(email); i++ {
		result = append(result, email[i])
	}

	return string(result)
}
