package main

func countStudents(students []int, sandwiches []int) int {
	var studentIndex, sandwichIndex, i int

	for {
		for i, studentIndex = 0, 0; i < len(students); i++ {
			if students[i] == sandwiches[sandwichIndex] {
				sandwichIndex++
			} else {
				students[studentIndex] = students[i]
				studentIndex++
			}
		}

		if studentIndex == len(students) {
			break
		}

		students = students[:studentIndex]
	}

	return len(students)
}
