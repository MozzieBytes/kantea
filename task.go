package main

type Status int

const (
	todo Status = iota
	inProgress
	done
)

type Task struct {
	title       string
	status      Status
	description string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}
