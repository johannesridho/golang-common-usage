// implement queue using doubly linked list

package main

import "container/list"

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{
		l: list.New(),
	}
}

func (q *Queue) Peek() interface{} {
	if q.l.Front() == nil {
		return nil
	}

	return q.l.Front().Value
}

func (q *Queue) Pop() interface{} {
	v := q.l.Front()

	if v == nil {
		return nil
	}

	q.l.Remove(v)
	return v.Value
}

func (q *Queue) Push(v interface{}) {
	q.l.PushBack(v)
}
