package main

func squareIsWhite(coordinates string) bool {
	first, second := coordinates[0]-'a', coordinates[1]-'1'
	return first&1 != second&1
}
