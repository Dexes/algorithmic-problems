package platform

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
)

type timus struct {
	taskTitles map[string]string
}

type timusTask struct {
	identifier string
	solutions  []string
}

func NewTimus() *timus {
	result := &timus{}

	file, _ := os.ReadFile("./_cache.timus.json")
	_ = json.Unmarshal(file, &result.taskTitles)

	return result
}

func (t *timus) GenerateReadme() {
	tasks := t.loadTasks()
	volumes := make([]string, 0, len(tasks))
	for k := range tasks {
		volumes = append(volumes, k)
	}
	sort.Strings(volumes)

	readmeContent := ""
	for _, volume := range volumes {
		readmeContent += t.generateVolumeTitle(volume) + "\n\n"
		for _, task := range tasks[volume] {
			readmeContent += t.generateTaskInfo(task) + "\n"
		}

		readmeContent += "\n"
	}

	file, _ := os.OpenFile("../timus/README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	_, _ = file.Write([]byte(strings.Trim(readmeContent, "\n") + "\n"))
	_ = file.Close()
}

func (t *timus) generateTaskInfo(task *timusTask) string {
	id := task.identifier
	title := t.taskTitles[id]

	result := "- [" + id + ". " + title + "](https://acm.timus.ru/problem.aspx?num=" + id + ")\n  "
	for _, solution := range task.solutions {
		language := strings.Split(solution, "/")[2]
		path := solution + "/main." + getLanguageExtension(language)
		result += "| [" + getLanguageTitle(language) + "](" + pathEscape(path) + ") "
	}

	return result[:len(result)-1]
}

func (t *timus) generateVolumeTitle(volume string) string {
	volumeNumber, _ := strconv.Atoi(volume[len(volume)-2:])
	start := strconv.Itoa(1000 + (volumeNumber-1)*100)
	end := strconv.Itoa(999 + volumeNumber*100)
	if volumeNumber == 12 {
		end = "2149"
	}

	return "## Volume " + strconv.Itoa(volumeNumber) + " (" + start + "-" + end + ")"

}

func (t *timus) loadTasks() map[string][]*timusTask {
	result := make(map[string][]*timusTask, 0)
	volumes, _ := os.ReadDir("../timus")
	for _, volume := range volumes {
		if !volume.IsDir() {
			continue
		}

		tasks, _ := os.ReadDir("../timus/" + volume.Name())
		for _, task := range tasks {
			languages, _ := os.ReadDir("../timus/" + volume.Name() + "/" + task.Name())
			solutions := make([]string, 0, len(languages))
			for _, language := range languages {
				solutions = append(solutions, volume.Name()+"/"+task.Name()+"/"+language.Name())
			}

			result[volume.Name()] = append(result[volume.Name()], &timusTask{
				identifier: task.Name(),
				solutions:  solutions,
			})
		}
	}

	return result
}
