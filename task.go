package main

type Status int

const (
	todo Status = iota
	inProgress
	done
)

func (s Status) Wrap() Status {
	switch {
	case s > done:
		s = todo
	case s < todo:
		s = done
	}
	return s
}

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
