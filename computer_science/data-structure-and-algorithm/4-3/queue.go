package main

import "fmt"

const ringBufferLength = 100

// queue リングバッファとして実装されたqueue
type queue struct {
	V    []task
	Head int // 先頭のインデックス
	Tail int // 次に要素が格納されるインデックス
}

func Newqueue(tasks []task) *queue {
	v := make([]task, ringBufferLength)
	copy(v, tasks)

	return &queue{
		V:    v,
		Head: 0,
		Tail: len(tasks),
	}
}

func (q *queue) isEmpty() bool {
	return q.Head == q.Tail
}

func (q *queue) enqueue(i task) {
	fmt.Printf("tail: %d", q.Tail)
	q.V[q.Tail] = i
	q.Tail = (q.Tail + 1) % len(q.V)
}

func (q *queue) dequeue() task {
	first := q.V[q.Head]
	q.Head = (q.Head + 1) % len(q.V)
	return first
}

func (q *queue) String() string {
	return fmt.Sprintf(":%v", q.V)
}
