package main

const (
	bulky   = "Bulky"
	heavy   = "Heavy"
	both    = "Both"
	neither = "Neither"
)

func categorizeBox(length int, width int, height int, mass int) string {
	volume := length * width * height
	isBulky := length >= 1e4 || width >= 1e4 || height >= 1e4 || volume >= 1e9
	isHeavy := mass >= 1e2

	switch {
	case isBulky && isHeavy:
		return both

	case isBulky:
		return bulky

	case isHeavy:
		return heavy

	default:
		return neither
	}
}
