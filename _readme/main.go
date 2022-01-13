package main

import "readme/platform"

func main() {
	platform.NewTimus().GenerateReadme()
	platform.NewLeetCode().GenerateReadme()
}
