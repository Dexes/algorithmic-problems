package main

type LogSystem struct {
	records []Record // here could be RBTree
}

func Constructor() LogSystem {
	return LogSystem{records: make([]Record, 0, 128)}
}

func (x *LogSystem) Put(id int, time string) {
	x.records = append(x.records, Record{ID: id, Seconds: toSeconds(time, GranularitySecond)})
}

func (x *LogSystem) Retrieve(startStr, endStr, granularityStr string) []int {
	granularity := toGranularity(granularityStr)
	start, end := toSeconds(startStr, granularity), toSeconds(endStr, granularity)+granularity
	result := make([]int, 0, 64)

	for _, record := range x.records {
		if start <= record.Seconds && record.Seconds < end {
			result = append(result, record.ID)
		}
	}

	return result
}

func toSeconds(time string, granularity int) int {
	result := (int(time[0]-'0')*1000 + int(time[1]-'0')*100 + int(time[2]-'0')*10 + int(time[3]-'0')) * GranularityYear

	timeIndex := 5
	for i := 1; i < len(gList) && gList[i] >= granularity; i++ {
		result += (int(time[timeIndex]-'0')*10 + int(time[timeIndex+1]-'0')) * gList[i]
		timeIndex += 3
	}

	return result
}

func toGranularity(s string) int {
	switch {
	case s[0] == 'Y':
		return GranularityYear

	case s[0] == 'M':
		if s[1] == 'o' {
			return GranularityMonth
		}

		return GranularityMinute

	case s[0] == 'D':
		return GranularityDay
	case s[0] == 'H':
		return GranularityHour
	default:
		return GranularitySecond
	}
}

type Record struct {
	ID      int
	Seconds int
}

const (
	GranularityYear   = 980294400
	GranularityMonth  = 2678400
	GranularityDay    = 86400
	GranularityHour   = 3600
	GranularityMinute = 60
	GranularitySecond = 1
)

var gList = []int{GranularityYear, GranularityMonth, GranularityDay, GranularityHour, GranularityMinute, GranularitySecond}
