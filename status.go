package main

type Status int

const (
	todo Status = iota
	inProgress
	done
)

func (s Status) Progress() Status {
	s++
	return s.wrap()
}

func (s Status) Regress() Status {
	s--
	return s.wrap()
}

func (s Status) wrap() Status {
	switch {
	case s > done:
		return todo
	case s < todo:
		return done
	default:
		return s
	}
}
