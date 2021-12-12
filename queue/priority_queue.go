package queue

import (
	"container/heap"
	"sync"
)

type LessFunc func(i interface{}, j interface{}) bool

type priorityElements struct {
	LessFunc LessFunc
	elements []interface{}
}

func (pe priorityElements) Len() int {
	return len(pe.elements)
}

func (pe priorityElements) Less(i, j int) bool {
	return pe.LessFunc(pe.elements[i], pe.elements[j])
}

func (pe priorityElements) Swap(i, j int) {
	pe.elements[i], pe.elements[j] = pe.elements[j], pe.elements[i]
}

func (pe *priorityElements) Push(e interface{}) {
	pe.elements = append(pe.elements, e)
}

func (pe *priorityElements) Pop() interface{} {
	old := pe.elements
	n := len(old)
	element := old[n-1]
	old[n-1] = nil
	pe.elements = old[:n-1]

	return element
}

type priorityQueue struct {
	sync.Mutex
	queue *priorityElements
}

func NewPriorityQueue(lessFunc LessFunc) Queue {
	return &priorityQueue{
		queue: &priorityElements{
			LessFunc: lessFunc,
		},
	}
}

func (q *priorityQueue) Len() int {
	q.Lock()
	defer q.Unlock()

	return q.queue.Len()
}

func (q *priorityQueue) Empty() bool {
	return q.queue.Len() == 0
}

func (q *priorityQueue) Push(e interface{}) {
	q.Lock()
	defer q.Unlock()
	heap.Push(q.queue, e)
}

func (q *priorityQueue) Front() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()

	return q.queue.elements[0], true
}

func (q *priorityQueue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()

	return heap.Pop(q.queue), true
}
