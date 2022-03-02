package main

var months = map[string]string{
	"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06",
	"Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
}

func reformatDate(date string) string {
	result := make([]byte, 0, 10)
	result = extractYear(date, result)
	result = extractMonth(date, result)
	result = extractDay(date, result)

	return string(result)
}

func extractDay(date string, destination []byte) []byte {
	if '0' <= date[1] && date[1] <= '9' {
		return append(destination, '-', date[0], date[1])
	}

	return append(destination, '-', '0', date[0])
}

func extractMonth(date string, destination []byte) []byte {
	month := months[date[len(date)-8:len(date)-5]]
	return append(destination, '-', month[0], month[1])
}

func extractYear(date string, destination []byte) []byte {
	return append(destination, date[len(date)-4:]...)
}
