package main

func haveConflict(event1 []string, event2 []string) bool {
	if event1[0] > event2[0] {
		event1, event2 = event2, event1
	}

	return event1[1] >= event2[0]
}
