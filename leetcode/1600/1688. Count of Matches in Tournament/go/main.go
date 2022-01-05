package main

func numberOfMatches(teams int) int {
	matches := 0
	for teams > 1 {
		matches += teams >> 1
		teams = (teams + 1) >> 1
	}

	return matches
}
