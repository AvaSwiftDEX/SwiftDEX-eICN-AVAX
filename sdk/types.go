package sdk

import "container/list"

type Queue struct {
	items *list.List
}

// create a new queue
func NewQueue() *Queue {
	return &Queue{items: list.New()}
}

// enqueue
func (q *Queue) Enqueue(item interface{}) {
	q.items.PushBack(item)
}

// dequeue
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.items.Len() == 0 {
		return nil, false // queue is empty
	}
	front := q.items.Front()
	q.items.Remove(front) // remove the first element
	return front.Value, true
}

// get the size of the queue
func (q *Queue) Size() int {
	return q.items.Len()
}

// check if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.items.Len() == 0
}
