package queue

import "container/list"

// Queue definition using a list
type Queue struct {
	list.List
}

// Enqueue adds an element to the back of the queue
func (q *Queue) Enqueue(v interface{}) *list.Element {
	return q.PushBack(v)
}

// Dequeue removes an element from the front of the queue
func (q *Queue) Dequeue() *list.Element {
	elem := q.Front()
	if elem != nil {
		q.Remove(elem)
	}
	return elem
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
