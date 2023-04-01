package platform

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	difficultyEasy   = 0
	difficultyMedium = 1
	difficultyHard   = 2
)

type leetCode struct {
	taskInfo      map[string]taskInfo
	volumes       map[string]struct{}
	total, solved [3]int
}

type taskInfo struct {
	Name       string `json:"name"`
	Url        string `json:"url"`
	Difficulty int    `json:"difficulty"`
	Premium    bool   `json:"premium"`
}

type leetCodeTask struct {
	id        string
	solutions []string
}

func NewLeetCode() *leetCode {
	result := &leetCode{volumes: make(map[string]struct{})}

	file, _ := os.ReadFile("./_cache.leetcode.json")
	_ = json.Unmarshal(file, &result.taskInfo)

	result.countTotalTasks()
	return result
}

func (l *leetCode) GenerateReadme() {
	tasks := l.loadTasks()

	l.countSolvedTasks(tasks)
	l.generateReadme(tasks)
}

func (l *leetCode) generateReadme(data map[string][]*leetCodeTask) {
	readmeContent := strings.Builder{}

	for i := 0; ; i += 100 {
		volume := fmt.Sprintf("%04d", i)
		if _, exists := l.volumes[volume]; !exists {
			break
		}

		for _, task := range data[volume] {
			readmeContent.WriteString(l.generateTaskInfo(task))
			readmeContent.WriteByte('\n')
		}
	}

	header := "### easy " + fmt.Sprintf("%d / %d (%.2f%%)", l.solved[difficultyEasy], l.total[difficultyEasy], float64(l.solved[difficultyEasy])/float64(l.total[difficultyEasy])*100) +
		" &nbsp;|&nbsp; medium " + fmt.Sprintf("%d / %d (%.2f%%)", l.solved[difficultyMedium], l.total[difficultyMedium], float64(l.solved[difficultyMedium])/float64(l.total[difficultyMedium])*100) +
		" &nbsp;|&nbsp; hard " + fmt.Sprintf("%d / %d (%.2f%%)", l.solved[difficultyHard], l.total[difficultyHard], float64(l.solved[difficultyHard])/float64(l.total[difficultyHard])*100) + "\n\n"

	file, _ := os.OpenFile("../leetcode/README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	_, _ = file.Write([]byte(header + readmeContent.String()))
	_ = file.Close()
}

func (l *leetCode) generateTaskInfo(task *leetCodeTask) string {
	info := l.taskInfo[task.id]
	result := "- [" + info.Name + " (" + l.difficulty(info.Difficulty) + ")](" + info.Url + ")\n"
	for _, solution := range task.solutions {
		language := strings.Split(solution, "/")[2]
		path := solution + "/main." + getLanguageExtension(language)
		result += "  &nbsp;|&nbsp; [" + getLanguageTitle(language) + "](" + pathEscape(path) + ") "
	}

	return result[:len(result)-1]
}

func (l *leetCode) loadTasks() map[string][]*leetCodeTask {
	result := make(map[string][]*leetCodeTask)
	volumes, _ := os.ReadDir("../leetcode")
	for _, volume := range volumes {
		if !volume.IsDir() {
			continue
		}

		l.volumes[volume.Name()] = struct{}{}
		result[volume.Name()] = make([]*leetCodeTask, 0, 100)

		tasks, _ := os.ReadDir("../leetcode/" + volume.Name())
		for _, task := range tasks {
			languages, _ := os.ReadDir("../leetcode/" + volume.Name() + "/" + task.Name())
			solutions := make([]string, 0, len(languages))
			for _, language := range languages {
				solutions = append(solutions, volume.Name()+"/"+task.Name()+"/"+language.Name())
			}

			result[volume.Name()] = append(result[volume.Name()], &leetCodeTask{
				id:        task.Name(),
				solutions: solutions,
			})
		}
	}

	return result
}

func (l *leetCode) countSolvedTasks(data map[string][]*leetCodeTask) {
	for _, tasks := range data {
		for i := 0; i < len(tasks); i++ {
			if info, ok := l.taskInfo[tasks[i].id]; ok {
				l.solved[info.Difficulty]++
				continue
			}

			panic(`can't find task info: "` + tasks[i].id + `"`)
		}
	}
}

func (l *leetCode) countTotalTasks() {
	for _, data := range l.taskInfo {
		l.total[data.Difficulty]++
	}
}

func (l *leetCode) difficulty(d int) string {
	switch d {
	case 0:
		return "easy"
	case 1:
		return "medium"
	default:
		return "hard"
	}
}
