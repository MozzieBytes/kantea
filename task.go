package main

type Task struct {
	title       string
	status      Status
	description string
}

func (t *Task) Next() Task {
	t.status = t.status.Progress()
	return *t
}

func (t *Task) Previous() Task {
	t.status = t.status.Regress()
	return *t
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
