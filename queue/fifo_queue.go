package queue

import (
	"container/list"
	"sync"
)

type fifoQueue struct {
	sync.Mutex
	list *list.List
}

func NewFifoQueue() Queue {
	return &fifoQueue{
		list: list.New(),
	}
}

func (q *fifoQueue) Len() int {
	q.Lock()
	defer q.Unlock()

	return q.list.Len()
}

func (q *fifoQueue) Empty() bool {
	return q.Len() == 0
}

func (q *fifoQueue) Push(elem interface{}) {
	q.Lock()
	defer q.Unlock()
	q.list.PushBack(elem)
}

func (q *fifoQueue) Front() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()

	if q.list.Len() == 0 {
		return nil, false
	}

	return q.list.Front(), true
}

func (q *fifoQueue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()

	if q.list.Len() == 0 {
		return nil, false
	}

	e := q.list.Front()

	return q.list.Remove(e), true
}
