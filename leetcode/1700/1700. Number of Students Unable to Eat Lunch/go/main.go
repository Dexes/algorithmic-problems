package main

func countStudents(students []int, sandwiches []int) int {
	var student, x, i int
	sandwichIndex, popIndex, pushIndex, queueLength := 0, 0, len(students), len(students)
	students = append(students, students...)

	for x != queueLength {
		for i, x = 0, queueLength; i < x; i++ {
			student, popIndex = queuePop(students, popIndex)
			if student == sandwiches[sandwichIndex] {
				sandwichIndex++
				queueLength--
			} else {
				pushIndex = queuePush(students, student, pushIndex)
			}
		}
	}

	return queueLength
}

func queuePop(queue []int, popIndex int) (int, int) {
	result := queue[popIndex]
	popIndex++

	if popIndex == len(queue) {
		return result, 0
	}

	return result, popIndex
}

func queuePush(queue []int, element, pushIndex int) int {
	queue[pushIndex] = element
	pushIndex++

	if pushIndex == len(queue) {
		return 0
	}

	return pushIndex
}
