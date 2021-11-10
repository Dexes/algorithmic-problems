package main

import "./platform"

func main() {
	platform.NewTimus().GenerateReadme()
	platform.NewLeetCode().GenerateReadme()
}
