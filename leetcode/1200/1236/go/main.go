package main

type HtmlParser struct {
}

func (p HtmlParser) GetUrls(_ string) []string { return nil }

type set map[string]struct{}

var buffer = make([]string, 0, 256)

func crawl(url string, parser HtmlParser) []string {
	return _crawl(url, hostname(url), parser, make(set), buffer)
}

func _crawl(url, host string, parser HtmlParser, visited set, result []string) []string {
	if _, ok := visited[url]; ok {
		return result
	}

	visited[url] = struct{}{}
	result = append(result, url)

	for _, sub := range parser.GetUrls(url) {
		subhost := hostname(sub)
		if host != subhost {
			continue
		}

		result = _crawl(sub, subhost, parser, visited, result)
	}

	return result
}

func hostname(url string) string {
	i, j := 7, 8
	if url[4] == 's' {
		i, j = 8, 9
	}

	for ; j < len(url) && url[j] != '/' && url[j] != ':'; j++ {
	}

	return url[i:j]
}
