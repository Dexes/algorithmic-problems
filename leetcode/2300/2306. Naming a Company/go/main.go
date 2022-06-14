package main

func distinctNames(ideas []string) int64 {
	result, suffixes := 0, makeSuffixes(ideas)

	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			intersections := countIntersections(suffixes[i], suffixes[j])
			result += 2 * (len(suffixes[i]) - intersections) * (len(suffixes[j]) - intersections)
		}
	}

	return int64(result)
}

func makeSuffixes(ideas []string) []map[string]struct{} {
	var set map[string]struct{}
	suffixes := make([]map[string]struct{}, 26)

	for i := 0; i < len(ideas); i++ {
		if set = suffixes[ideas[i][0]-'a']; set == nil {
			set = make(map[string]struct{})
			suffixes[ideas[i][0]-'a'] = set
		}

		set[ideas[i][1:]] = struct{}{}
	}

	return suffixes
}

func countIntersections(a, b map[string]struct{}) (result int) {
	for suffix := range a {
		if _, exists := b[suffix]; exists {
			result++
		}
	}

	return result
}
