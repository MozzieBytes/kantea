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
