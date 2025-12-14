package main

func main_(
	q int,
	input []task,
) []task {
	queue := Newqueue(input)
	var totalProgressTime int
	doneSlice := make([]task, 0, len(input))
	for {
		if queue.isEmpty() {
			break
		}
		if queue.isEmpty() {
			return doneSlice
		}
		first := queue.dequeue()
		executionTime := min(q, first.Time) // ここが思いつかなかった
		first.process(executionTime)
		totalProgressTime += executionTime
		if !first.hasDone() {
			queue.enqueue(first)
		} else {
			first.setDoneTime(totalProgressTime)
			doneSlice = append(doneSlice, first)
		}
	}
	return doneSlice
}
