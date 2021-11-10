package platform

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type leetCode struct {
	taskUrls map[string]string
}

type leetCodeTask struct {
	name      string
	solutions []string
}

func NewLeetCode() *leetCode {
	result := &leetCode{}

	file, _ := ioutil.ReadFile("./_cache.leetcode.urls.json")
	json.Unmarshal(file, &result.taskUrls)

	return result
}

func (l *leetCode) GenerateReadme() {
	readmeContent := ""
	for _, task := range l.loadTasks() {
		readmeContent += l.generateTaskInfo(task) + "\n"
	}

	file, _ := os.OpenFile("../leetcode/README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	file.Write([]byte(strings.Trim(readmeContent, "\n") + "\n"))
	file.Close()
}

func (l *leetCode) generateTaskInfo(task *leetCodeTask) string {
	result := "- [" + task.name + "](" + l.taskUrls[task.name] + ")\n  "
	for _, solution := range task.solutions {
		language := strings.Split(solution, "/")[2]
		path := solution + "/main." + getLanguageExtension(language)
		result += "| [" + getLanguageTitle(language) + "](" + strings.ReplaceAll(path, " ", "%20") + ") "
	}

	return result[:len(result)-1]
}

func (l *leetCode) loadTasks() []*leetCodeTask {
	result := make([]*leetCodeTask, 0)
	volumes, _ := ioutil.ReadDir("../leetcode")
	for _, volume := range volumes {
		if !volume.IsDir() {
			continue
		}

		tasks, _ := ioutil.ReadDir("../leetcode/" + volume.Name())
		for _, task := range tasks {
			languages, _ := ioutil.ReadDir("../leetcode/" + volume.Name() + "/" + task.Name())
			solutions := make([]string, 0, len(languages))
			for _, language := range languages {
				solutions = append(solutions, volume.Name()+"/"+task.Name()+"/"+language.Name())
			}

			result = append(result, &leetCodeTask{
				name:      strings.TrimLeft(task.Name(), "0"),
				solutions: solutions,
			})
		}
	}

	return result
}
