package main

func mostPopularCreator(creators []string, videos []string, views []int) [][]string {
	data, max := make(map[string]*CreatorInfo), 0
	for i := 0; i < len(creators); i++ {
		info := updateInfo(data[creators[i]], videos[i], views[i])
		if info.CreatorViews > max {
			max = info.CreatorViews
		}

		data[creators[i]] = info
	}

	result := make([][]string, 0, 5)
	for name, info := range data {
		if info.CreatorViews == max {
			result = append(result, []string{name, info.Video})
		}
	}

	return result
}

func updateInfo(info *CreatorInfo, video string, views int) *CreatorInfo {
	if info == nil {
		return &CreatorInfo{Video: video, VideoViews: views, CreatorViews: views}
	}

	if info.VideoViews < views || (info.VideoViews == views && video < info.Video) {
		info.Video, info.VideoViews = video, views
	}

	info.CreatorViews += views
	return info
}

type CreatorInfo struct {
	Video        string
	VideoViews   int
	CreatorViews int
}
