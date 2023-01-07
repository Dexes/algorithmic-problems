package kata

func ArrayDiff(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	set := make(map[int]struct{})
	for _, n := range b {
		set[n] = struct{}{}
	}

	wIndex := 0
	for i := 0; i < len(a); i++ {
		if _, ok := set[a[i]]; !ok {
			a[wIndex], wIndex = a[i], wIndex+1
		}
	}

	return a[:wIndex]
}
