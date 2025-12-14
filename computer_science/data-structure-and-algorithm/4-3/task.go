package main

type task struct {
	Name string
	Time int
}

func (i *task) process(q int) {
	i.Time -= q
}

func (i *task) hasDone() bool {
	return i.Time <= 0
}

func (i *task) setDoneTime(totaProgresslTime int) {
	i.Time = totaProgresslTime
}
