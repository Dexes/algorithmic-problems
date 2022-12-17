package main

type AuthenticationManager struct {
	tokens map[string]int
	ttl    int
}

func Constructor(ttl int) AuthenticationManager {
	return AuthenticationManager{tokens: make(map[string]int), ttl: ttl}
}

func (x AuthenticationManager) Generate(tokenId string, currentTime int) {
	x.tokens[tokenId] = currentTime + x.ttl
}

func (x AuthenticationManager) Renew(tokenId string, currentTime int) {
	if x.tokens[tokenId] > currentTime {
		x.Generate(tokenId, currentTime)
	}
}

func (x AuthenticationManager) CountUnexpiredTokens(currentTime int) (result int) {
	for _, time := range x.tokens {
		if time > currentTime {
			result++
		}
	}

	return result
}
