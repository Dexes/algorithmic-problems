package main

import (
	"sort"
)

var buffer = make([]string, 0, 100)

type TodoList struct {
	lastID int
	tasks  [101][]task
}

type task struct {
	id          int
	userID      int
	dueDate     int
	description string
	tags        map[string]struct{}
}

func Constructor() TodoList {
	return TodoList{}
}

func (x *TodoList) AddTask(userId int, taskDescription string, dueDate int, tags []string) int {
	x.lastID++
	x.tasks[userId] = insert(
		x.tasks[userId],
		task{id: x.lastID, userID: userId, dueDate: dueDate, description: taskDescription, tags: toSet(tags)},
	)

	return x.lastID
}

func (x *TodoList) GetAllTasks(userId int) []string {
	result := buffer
	for _, t := range x.tasks[userId] {
		result = append(result, t.description)
	}

	return result
}

func (x *TodoList) GetTasksForTag(userId int, tag string) []string {
	result := buffer[:0]
	for _, t := range x.tasks[userId] {
		if _, ok := t.tags[tag]; !ok {
			continue
		}

		result = append(result, t.description)
	}

	return result
}

func (x *TodoList) CompleteTask(userId int, taskId int) {
	x.tasks[userId] = remove(x.tasks[userId], taskId)
}

func remove(tasks []task, id int) []task {
	for i, t := range tasks {
		if t.id != id {
			continue
		}

		copy(tasks[i:], tasks[i+1:])
		return tasks[:len(tasks)-1]
	}

	return tasks
}

func insert(tasks []task, t task) []task {
	if tasks == nil {
		return append(make([]task, 0, 8), t)
	}

	if len(tasks) == 0 || tasks[len(tasks)-1].dueDate < t.dueDate {
		return append(tasks, t)
	}

	index := sort.Search(len(tasks), func(i int) bool { return tasks[i].dueDate > t.dueDate })

	tasks = append(tasks, task{})
	copy(tasks[index+1:], tasks[index:])
	tasks[index] = t

	return tasks
}

func toSet(x []string) map[string]struct{} {
	result := make(map[string]struct{})
	for i := 0; i < len(x); i++ {
		result[x[i]] = struct{}{}
	}

	return result
}
