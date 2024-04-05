package main

import "container/list"

// we may need to use a queue for optimizations later
type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Enqueue(value any) {
	q.list.PushBack(value)
}

func (q *Queue) Dequeue() any {
	front := q.list.Front()
	if front != nil {
		q.list.Remove(front)
		return front.Value
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}
