package main

type BrowserHistory struct {
	history      []string
	historyIndex int
}

func Constructor(homepage string) BrowserHistory {
	history, historyIndex := make([]string, 1, 500), 0
	history[historyIndex] = homepage

	return BrowserHistory{history: history, historyIndex: 0}
}

func (x *BrowserHistory) Visit(url string) {
	x.historyIndex++
	x.history = append(x.history[:x.historyIndex], url)
}

func (x *BrowserHistory) Back(steps int) string {
	x.historyIndex = max(x.historyIndex-steps, 0)
	return x.history[x.historyIndex]
}

func (x *BrowserHistory) Forward(steps int) string {
	x.historyIndex = min(x.historyIndex+steps, len(x.history)-1)
	return x.history[x.historyIndex]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
