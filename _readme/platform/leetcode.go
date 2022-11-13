package platform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	difficultyEasy   = 0
	difficultyMedium = 1
	difficultyHard   = 2
)

type leetCode struct {
	taskInfo map[string]struct {
		Url        string `json:"url"`
		Difficulty int    `json:"difficulty"`
		Premium    bool   `json:"premium"`
	}
	volumes                   map[string]struct{}
	easyTotal, easySolved     map[string]int
	mediumTotal, mediumSolved map[string]int
	hardTotal, hardSolved     map[string]int
}

type leetCodeTask struct {
	name      string
	solutions []string
}

func NewLeetCode() *leetCode {
	result := &leetCode{
		volumes:      make(map[string]struct{}),
		easyTotal:    make(map[string]int),
		easySolved:   make(map[string]int),
		mediumTotal:  make(map[string]int),
		mediumSolved: make(map[string]int),
		hardTotal:    make(map[string]int),
		hardSolved:   make(map[string]int),
	}

	file, _ := ioutil.ReadFile("./_cache.leetcode.urls.json")
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
	easySolved, easyTotal, mediumSolved, mediumTotal, hardSolved, hardTotal := 0, 0, 0, 0, 0, 0

	for i := 0; ; i += 100 {
		volume := fmt.Sprintf("%04d", i)
		if _, exists := l.volumes[volume]; !exists {
			break
		}

		easySolved, easyTotal = easySolved+l.easySolved[volume], easyTotal+l.easyTotal[volume]
		mediumSolved, mediumTotal = mediumSolved+l.mediumSolved[volume], mediumTotal+l.mediumTotal[volume]
		hardSolved, hardTotal = hardSolved+l.hardSolved[volume], hardTotal+l.hardTotal[volume]

		//vMax := []byte(volume)
		//vMax[2], vMax[3] = '9', '9'
		//
		//readmeContent.WriteString("<details><summary>" + volume + " - " + string(vMax))
		//readmeContent.WriteString(" &nbsp;|&nbsp; easy " + strconv.Itoa(l.easySolved[volume]) + " / " + strconv.Itoa(l.easyTotal[volume]))
		//readmeContent.WriteString(" &nbsp;|&nbsp; medium " + strconv.Itoa(l.mediumSolved[volume]) + " / " + strconv.Itoa(l.mediumTotal[volume]))
		//readmeContent.WriteString(" &nbsp;|&nbsp; hard " + strconv.Itoa(l.hardSolved[volume]) + " / " + strconv.Itoa(l.hardTotal[volume]))
		//readmeContent.WriteString("</summary>\n\n")

		for _, task := range data[volume] {
			readmeContent.WriteString(l.generateTaskInfo(task))
			readmeContent.WriteByte('\n')
		}

		//readmeContent.WriteString("</details>\n")
	}

	header := "### easy " + fmt.Sprintf("%d / %d (%.2f%%)", easySolved, easyTotal, float64(easySolved)/float64(easyTotal)*100) +
		" &nbsp;|&nbsp; medium " + fmt.Sprintf("%d / %d (%.2f%%)", mediumSolved, mediumTotal, float64(mediumSolved)/float64(mediumTotal)*100) +
		" &nbsp;|&nbsp; hard " + fmt.Sprintf("%d / %d (%.2f%%)", hardSolved, hardTotal, float64(hardSolved)/float64(hardTotal)*100) + "\n\n"

	file, _ := os.OpenFile("../leetcode/README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	_, _ = file.Write([]byte(header + readmeContent.String()))
	_ = file.Close()
}

func (l *leetCode) generateTaskInfo(task *leetCodeTask) string {
	info := l.taskInfo[task.name]
	result := "- [" + task.name + " (" + l.difficulty(info.Difficulty) + ")](" + info.Url + ")\n"
	for _, solution := range task.solutions {
		language := strings.Split(solution, "/")[2]
		path := solution + "/main." + getLanguageExtension(language)
		result += "  &nbsp;|&nbsp; [" + getLanguageTitle(language) + "](" + pathEscape(path) + ") "
	}

	return result[:len(result)-1]
}

func (l *leetCode) loadTasks() map[string][]*leetCodeTask {
	result := make(map[string][]*leetCodeTask)
	volumes, _ := ioutil.ReadDir("../leetcode")
	for _, volume := range volumes {
		if !volume.IsDir() {
			continue
		}

		result[volume.Name()] = make([]*leetCodeTask, 0, 100)

		tasks, _ := ioutil.ReadDir("../leetcode/" + volume.Name())
		for _, task := range tasks {
			languages, _ := ioutil.ReadDir("../leetcode/" + volume.Name() + "/" + task.Name())
			solutions := make([]string, 0, len(languages))
			for _, language := range languages {
				solutions = append(solutions, volume.Name()+"/"+task.Name()+"/"+language.Name())
			}

			result[volume.Name()] = append(result[volume.Name()], &leetCodeTask{
				name:      strings.TrimLeft(strings.ReplaceAll(task.Name(), "❔", "?"), "0"),
				solutions: solutions,
			})
		}
	}

	return result
}

func (l *leetCode) countSolvedTasks(data map[string][]*leetCodeTask) {
	for volume, tasks := range data {
		for i := 0; i < len(tasks); i++ {
			switch info, ok := l.taskInfo[tasks[i].name]; {
			case !ok:
				panic(`can't find task info: "` + tasks[i].name + `"`)
			case info.Difficulty == difficultyEasy:
				l.easySolved[volume]++
			case info.Difficulty == difficultyMedium:
				l.mediumSolved[volume]++
			case info.Difficulty == difficultyHard:
				l.hardSolved[volume]++
			default:
				panic("difficulty " + strconv.Itoa(info.Difficulty))
			}
		}
	}
}

func (l *leetCode) countTotalTasks() {
	for title, data := range l.taskInfo {
		volume := l.getVolume(title)
		l.volumes[volume] = struct{}{}

		switch data.Difficulty {
		case difficultyEasy:
			l.easyTotal[volume]++
		case difficultyMedium:
			l.mediumTotal[volume]++
		case difficultyHard:
			l.hardTotal[volume]++
		default:
			panic("difficulty " + strconv.Itoa(data.Difficulty))
		}
	}
}

func (l *leetCode) getVolume(title string) string {
	x := 0
	for i := 0; title[i] != '.'; i++ {
		x = x*10 + int(title[i]-'0')
	}

	return fmt.Sprintf("%04d", x/100*100)
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
