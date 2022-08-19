package main

type MyCalendar struct {
	events *Event
}

func Constructor() MyCalendar {
	return MyCalendar{events: &Event{Start: -1, End: -1}}
}

func (x *MyCalendar) Book(start int, end int) bool {
	prev, current := x.events, x.events.Next
	for ; ; prev, current = current, current.Next {
		if current != nil && end > current.Start {
			continue
		}

		if prev.End > start {
			return false
		}

		prev.Next = &Event{Start: start, End: end, Next: current}
		return true
	}
}

type Event struct {
	Start int
	End   int
	Next  *Event
}
