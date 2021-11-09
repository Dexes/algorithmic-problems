package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var titles = make(map[string]string)

func main() {
	tasks := getTasksList()
	volumes := make(map[string][]string)

	for _, id := range getSortedKeys(tasks) {
		appendToVolumes(volumes, id, tasks[id])
	}

	var readme = ""
	var index = 0
	var re = regexp.MustCompile(`0+(\d+)`)
	for _, volume := range getSortedKeys(volumes) {
		var start = strconv.Itoa(1000 + index*100)
		var end = strconv.Itoa(999 + (index+1)*100)
		if index == 11 {
			end = "2149"
		}
		var volumeTitle = string(re.ReplaceAll([]byte(volume), []byte("$1")))
		readme += "## " + volumeTitle + " (" + start + "-" + end + ")\n"
		readme += strings.Join(volumes[volume], "\n") + "\n\n"
		index++
	}

	file, _ := os.OpenFile("../README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	file.Write([]byte(strings.Trim(readme, "\n") + "\n"))
	file.Close()
}

func getTasksList() map[string][]string {
	var startPath = "../"
	var result = make(map[string][]string)

	filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		path = strings.ReplaceAll(path, "\\", "/")
		path = strings.TrimLeft(path, startPath)
		if !strings.HasPrefix(path, "Volume") || strings.Contains(path, "test") {
			return nil
		}

		var id = strings.Split(path, "/")[1]
		if _, ok := result[id]; !ok {
			result[id] = make([]string, 0)
		}

		result[id] = append(result[id], path)

		return nil
	})

	return result
}

func appendToVolumes(volumes map[string][]string, id string, task []string) {
	var title = getTaskTitle(id)
	var volume = strings.Split(task[0], "/")[0]
	var value = "- [" + id + ". " + title + "](https://acm.timus.ru/problem.aspx?num=" + id + ")\n"
	for _, path := range task {
		var data = strings.Split(path, "/")
		value += "| [" + getLanguageTitle(data[2]) + "](" + strings.ReplaceAll(path, " ", "%20") + ") "
	}

	volumes[volume] = append(volumes[volume], strings.Trim(value, " "))
}

func getLanguageTitle(lang string) string {
	switch lang {
	case "cpp":
		return "C++"
	case "csharp":
		return "C#"
	default:
		return lang
	}
}

func getTaskTitle(id string) string {
	if len(titles) == 0 {
		file, _ := ioutil.ReadFile("titles.json")
		_ = json.Unmarshal(file, &titles)
	}

	return titles[id]
}

func getSortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
